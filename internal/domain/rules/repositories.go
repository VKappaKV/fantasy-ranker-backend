package rules

import (
	"context"

	"github.com/VKappaKV/fantasy-ranker-backend/internal/domain/models"
)

type PlayerRepository interface {
	UpsertByPUUID(ctx context.Context, p models.Player) error
	ByPUUID(ctx context.Context, puuid string) (models.Player, error)
}
