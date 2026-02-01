package rules

import "github.com/VKappaKV/fantasy-ranker-backend/internal/domain/models"

func Ratio(k models.KDA) float64 {
	if k.Deaths == 0 {
		return float64(k.Kills + k.Assists)
	}
	return float64(k.Kills+k.Assists) / float64(k.Deaths)
}