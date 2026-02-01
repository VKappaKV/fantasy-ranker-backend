package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/VKappaKV/fantasy-ranker-backend/internal/domain/rules"
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
		region, err := rules.ParseRegion(r.URL.Query().Get("region"))
		if err != nil {
			writeAPIError(w, http.StatusBadRequest, "INVALID_REGION", err.Error(), nil)
			return
		}

		puuid := strings.TrimSpace(r.URL.Query().Get("puuid"))
		if puuid == "" {
			writeAPIError(w, http.StatusBadRequest, "MISSING_PARAM", "missing required query param: puuid", nil)
			return
		}

		start := parseIntDefault(r.URL.Query().Get("start"), 0)
		count := parseIntDefault(r.URL.Query().Get("count"), 20)

		if start < 0 {
			start = 0
		}
		if count <= 0 {
			count = 20
		}
		if count > 100 {
			count = 100
		}

		ids, err := c.MatchIDsByPUUID(r.Context(), string(region), puuid, start, count)
		if err != nil {
			writeMappedError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, riotMatchesResponse{
			PUUID:    puuid,
			Region:   string(region),
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
