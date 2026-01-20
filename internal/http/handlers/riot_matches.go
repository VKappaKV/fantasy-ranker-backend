package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/VKappaKV/fantasy-ranker-backend/internal/riot"
)

type riotMatchesResponse struct {
	PUUID    string   `json:"puuid"`
	Region   string   `json:"region"`
	Start    int      `json:"start"`
	Count    int      `json:"count"`
	MatchIDs []string `json:"matchIds"`
}

func RiotMatches(c *riot.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		region := strings.TrimSpace(r.URL.Query().Get("region"))
		puuid := strings.TrimSpace(r.URL.Query().Get("puuid"))

		if region == "" || puuid == "" {
			writeJSON(w, http.StatusBadRequest, errorResponse{
				Error: "missing required query params: region, puuid",
			})
			return
		}

		start := parseIntDefault(r.URL.Query().Get("start"), 0)
		count := parseIntDefault(r.URL.Query().Get("count"), 20)

		// clamp di sicurezza
		if start < 0 {
			start = 0
		}
		if count <= 0 {
			count = 20
		}
		if count > 100 {
			count = 100
		}

		ids, err := c.MatchIDsByPUUID(r.Context(), region, puuid, start, count)
		if err != nil {
			writeRiotError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, riotMatchesResponse{
			PUUID:    puuid,
			Region:   region,
			Start:    start,
			Count:    count,
			MatchIDs: ids,
		})
	}
}

func parseIntDefault(s string, def int) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return def
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return n
}
