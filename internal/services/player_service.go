package services

import (
	"context"

	d "github.com/VKappaKV/fantasy-ranker-backend/internal/domain"
	"github.com/VKappaKV/fantasy-ranker-backend/internal/riot"
)

type PlayerService struct {
	riot *riot.Client
	repo d.PlayerRepository
}

func NewPlayerService(riotClient *riot.Client, repo d.PlayerRepository) *PlayerService {
	return &PlayerService{
		riot: riotClient,
		repo: repo,
	}
}

func (s *PlayerService) RegisterPlayerByRiotID(ctx context.Context, region d.Region, riotID d.RiotID) (d.Player, error) {
	acc, err := s.riot.AccountByRiotID(ctx, string(region), riotID.GameName, riotID.TagLine)

	if err != nil {
		return d.Player{}, err
	}

	p := d.Player{
		PUUID: acc.PUUID,
		Region: region,
		GameName: acc.GameName,
		TagLine: acc.TagLine,
	}

	return s.repo.UpsertByPUUID(ctx, p)
}