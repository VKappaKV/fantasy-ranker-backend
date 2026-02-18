package tui

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Endpoint struct {
	Name        string
	Method      string
	Path        string
	Help 		string
	Params      []ParamSpec // query params for now
}

type ParamSpec struct {
	Key      string
	Required bool
	Hint     string
}

type APIResponse struct {
	Status     string
	StatusCode int
	Headers    http.Header
	BodyRaw    []byte
	BodyPretty string
	Duration   time.Duration
}

type Client struct {
	BaseURL string
	HTTP    *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: strings.TrimRight(baseURL, "/"),
		HTTP: &http.Client{
			Timeout: 20 * time.Second,
		},
	}
}

func (c *Client) Do(ctx context.Context, method, path string, query map[string]string, body any) (*APIResponse, error) {
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}
	q := u.Query()

	// stable ordering makes it easier to read/debug
	keys := make([]string, 0, len(query))
	for k := range query {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := strings.TrimSpace(query[k])
		if v != "" {
			q.Set(k, v)
		}
	}
	u.RawQuery = q.Encode()

	var rdr io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		rdr = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), rdr)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	start := time.Now()
	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	dur := time.Since(start)

	pretty := prettifyJSON(raw)

	return &APIResponse{
		Status:     res.Status,
		StatusCode: res.StatusCode,
		Headers:    res.Header,
		BodyRaw:    raw,
		BodyPretty: pretty,
		Duration:   dur,
	}, nil
}

func prettifyJSON(raw []byte) string {
	trim := bytes.TrimSpace(raw)
	if len(trim) == 0 {
		return ""
	}

	// If it's JSON, indent it. Otherwise show as text.
	var js any
	if err := json.Unmarshal(trim, &js); err != nil {
		return string(raw)
	}
	out, err := json.MarshalIndent(js, "", "  ")
	if err != nil {
		return string(raw)
	}
	return string(out)
}

func (r *APIResponse) SummaryLine() string {
	return fmt.Sprintf("%s (%d) • %s", r.Status, r.StatusCode, r.Duration.Round(time.Millisecond))
}
