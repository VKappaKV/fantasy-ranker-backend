package handlers

import (
	"net/http"

	"github.com/VKappaKV/fantasy-ranker-backend/internal/domain/rules"
	"github.com/VKappaKV/fantasy-ranker-backend/internal/riot"
)

type riotAccountResponse struct {
	RiotID   string `json:"riotId"`
	Region   string `json:"region"`
	PUUID    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

func RiotAccount(c *riot.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		region, err := rules.ParseRegion(r.URL.Query().Get("region"))

		if err != nil {
			writeAPIError(w, http.StatusBadRequest, "INVALID_REGION", err.Error(), nil)
			return
		}
				
		riotID, err := rules.ParseRiotID(r.URL.Query().Get("riotId"))
		if err != nil {
			writeAPIError(w, http.StatusBadRequest, "INVALID_RIOT_ID", err.Error(), nil)
			return
		}

		acc, err := c.AccountByRiotID(r.Context(), string(region), riotID.GameName, riotID.TagLine)
		if err != nil {
			writeMappedError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, riotAccountResponse{
			RiotID:   riotID.GameName + "#" + riotID.TagLine,
			Region:   string(region),
			PUUID:    acc.PUUID,
			GameName: acc.GameName,
			TagLine:  acc.TagLine,
		})
	}
}


