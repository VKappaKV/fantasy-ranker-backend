package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/VKappaKV/fantasy-ranker-backend/internal/riot"
)

type riotAccountResponse struct {
	RiotID   string `json:"riotId"`
	Region   string `json:"region"`
	PUUID    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func RiotAccount(c *riot.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		region := strings.TrimSpace(r.URL.Query().Get("region"))
		riotID := strings.TrimSpace(r.URL.Query().Get("riotId"))

		if region == "" || riotID == "" {
			writeJSON(w, http.StatusBadRequest, errorResponse{
				Error: "missing required query params: region, riotId",
			})
			return
		}

		gameName, tagLine, ok := splitRiotID(riotID)
		if !ok {
			writeJSON(w, http.StatusBadRequest, errorResponse{
				Error: "invalid riotId format. Expected GameName#TAG",
			})
			return
		}

		acc, err := c.AccountByRiotID(r.Context(), region, gameName, tagLine)
		if err != nil {
			writeRiotError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, riotAccountResponse{
			RiotID:   gameName + "#" + tagLine,
			Region:   region,
			PUUID:    acc.PUUID,
			GameName: acc.GameName,
			TagLine:  acc.TagLine,
		})
	}
}

func splitRiotID(riotID string) (gameName, tagLine string, ok bool) {
	parts := strings.Split(riotID, "#")
	if len(parts) != 2 {
		return "", "", false
	}
	gameName = strings.TrimSpace(parts[0])
	tagLine = strings.TrimSpace(parts[1])
	if gameName == "" || tagLine == "" {
		return "", "", false
	}
	return gameName, tagLine, true
}

func writeRiotError(w http.ResponseWriter, err error) {
	var re *riot.Error
	if errors.As(err, &re) {
		writeJSON(w, re.HTTPStatus, errorResponse{Error: re.Message})
		return
	}
	writeJSON(w, http.StatusBadGateway, errorResponse{Error: "upstream error"})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
