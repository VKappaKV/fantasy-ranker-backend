package services

import (
	d "github.com/VKappaKV/fantasy-ranker-backend/internal/domain"
)

func ChampionFromRiot(name string) (d.Champion, error) {
	switch name {
	case "Ahri":
		return d.ChampionAhri, nil
	case "LeeSin":
		return d.ChampionLeeSin, nil
	case "MissFortune":
		return d.ChampionMissFortune, nil
	default:
		return "", d.ErrUnknownChampion
	}
}
