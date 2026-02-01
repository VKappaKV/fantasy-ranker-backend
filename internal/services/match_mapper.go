package services

import (
	"time"

	models "github.com/VKappaKV/fantasy-ranker-backend/internal/domain/models"
	"github.com/VKappaKV/fantasy-ranker-backend/internal/riot"
)

func MapRiotMatchToDomain(dto riot.RiotMatch) models.Match {
	players := make([]models.MatchPlayer, 0, len(dto.Info.Participants))

	for _, p := range dto.Info.Participants {
		players = append(players, models.MatchPlayer{
			PlayerID: models.PlayerID(p.PUUID),
			Champion: models.Champion(p.Champion),
			KDA: models.KDA{
				Kills:   p.Kills,
				Deaths:  p.Deaths,
				Assists: p.Assists,
			},
			Win: p.Win,
		})
	}

	return models.Match{
		ID:       dto.Metadata.MatchID,
		Duration: time.Duration(dto.Info.GameDuration) * time.Second,
		Queue:    models.QueueFromRiotID(dto.Info.QueueID),
		Players:  players,
	}
}