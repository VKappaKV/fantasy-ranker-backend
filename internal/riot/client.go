package riot

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
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
		var retryAfter *int
		// Riot rate limit responses often include Retry-After header (seconds)
		if h := resp.Header.Get("Retry-After"); h != "" {
			if n, err := strconv.Atoi(h); err == nil {
				retryAfter = &n
			}
		}
		return parseRiotError(resp.StatusCode, body, retryAfter)
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
	queue := 420 // Ranked Solo/Duo queue ID
	_type := "ranked" // Match type filter
	u := fmt.Sprintf("https://%s.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?queue=%d&type=%s&start=%d&count=%d",
		region,
		url.PathEscape(puuid),
		queue,
		_type,
		start,
		count,
	)

	var ids []string
	if err := c.doJSON(ctx, http.MethodGet, u, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

func (c *Client) MatchByID(ctx context.Context, region, matchID string) (RiotMatch, error) {
	u := fmt.Sprintf("https://%s.api.riotgames.com/lol/match/v5/matches/%s",
		region,
		url.PathEscape(matchID),
	)
	var match RiotMatch
	if err := c.doJSON(ctx, http.MethodGet, u, &match); err != nil {
		return RiotMatch{}, err
	}
	return match, nil
}