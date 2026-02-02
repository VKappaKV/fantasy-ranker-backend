package services

import (
	"time"

	d "github.com/VKappaKV/fantasy-ranker-backend/internal/domain"
	"github.com/VKappaKV/fantasy-ranker-backend/internal/riot"
)

func MapRiotMatchToDomain(dto riot.RiotMatch) d.Match {
	players := make([]d.MatchPlayer, 0, len(dto.Info.Participants))

	for _, p := range dto.Info.Participants {
		players = append(players, d.MatchPlayer{
			PlayerID: d.PlayerID(p.PUUID),
			Champion: d.Champion(p.Champion),
			KDA: d.KDA{
				Kills:   p.Kills,
				Deaths:  p.Deaths,
				Assists: p.Assists,
			},
			Win: p.Win,
		})
	}

	return d.Match{
		ID:       dto.Metadata.MatchID,
		Duration: time.Duration(dto.Info.GameDuration) * time.Second,
		Queue:    d.QueueFromRiotID(dto.Info.QueueID),
		Players:  players,
	}
}