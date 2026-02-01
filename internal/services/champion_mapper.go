package services

import (
	models "github.com/VKappaKV/fantasy-ranker-backend/internal/domain/models"
	rules "github.com/VKappaKV/fantasy-ranker-backend/internal/domain/rules"
)

func ChampionFromRiot(name string) (models.Champion, error) {
	switch name {
	case "Ahri":
		return models.ChampionAhri, nil
	case "LeeSin":
		return models.ChampionLeeSin, nil
	case "MissFortune":
		return models.ChampionMissFortune, nil
	default:
		return "", rules.ErrUnknownChampion
	}
}
