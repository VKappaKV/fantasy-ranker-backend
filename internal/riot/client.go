package riot

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	apiKey string
	http   *http.Client
}

func New(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) doJSON(ctx context.Context, method, fullURL string, out any) error {
	req, err := http.NewRequestWithContext(ctx, method, fullURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Riot-Token", c.apiKey)

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return parseRiotError(resp.StatusCode, body)
	}

	if out == nil {
		return nil
	}
	return json.Unmarshal(body, out)
}

func (c *Client) AccountByRiotID(ctx context.Context, region, gameName, tagLine string) (Account, error) {
	u := fmt.Sprintf("https://%s.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s",
		region,
		url.PathEscape(gameName),
		url.PathEscape(tagLine),
	)

	var acc Account
	if err := c.doJSON(ctx, http.MethodGet, u, &acc); err != nil {
		return Account{}, err
	}
	return acc, nil
}

func (c *Client) MatchIDsByPUUID(ctx context.Context, region, puuid string, start, count int) ([]string, error) {
	u := fmt.Sprintf("https://%s.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=%d&count=%d",
		region,
		url.PathEscape(puuid),
		start,
		count,
	)

	var ids []string
	if err := c.doJSON(ctx, http.MethodGet, u, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}
