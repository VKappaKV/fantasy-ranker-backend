package services

import (
	domain "github.com/VKappaKV/fantasy-ranker-backend/internal/domain"
	models "github.com/VKappaKV/fantasy-ranker-backend/internal/domain/models"
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
		return "", domain.ErrUnknownChampion
	}
}
