package handlers

import (
	"encoding/json"
	"net/http"

	d "github.com/VKappaKV/fantasy-ranker-backend/internal/domain"
	"github.com/VKappaKV/fantasy-ranker-backend/internal/services"
)

type registerPlayerRequest struct {
	Region string `json:"region"`
	RiotID string `json:"riotId"`
}

type registerPlayerResponse struct {
	Player d.Player `json:"player"`
}

func RegisterPlayer(svc *services.PlayerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req registerPlayerRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeAPIError(w, http.StatusBadRequest, "INVALID_JSON", "invalid JSON body", nil)
			return
		}

		region, err := d.ParseRegion(req.Region)
		if err != nil {
			writeAPIError(w, http.StatusBadRequest, "INVALID_REGION", "region must be one of: europe, americas, asia", nil)
			return
		}

		riotID, err := d.ParseRiotID(req.RiotID)
		if err != nil {
			writeAPIError(w, http.StatusBadRequest, "INVALID_RIOT_ID", "riotId must be in the format GameName#TAG", nil)
			return
		}

		player, err := svc.RegisterPlayerByRiotID(r.Context(), region, riotID)
		if err != nil {
			writeMappedError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, registerPlayerResponse{Player: player})
	}
}
