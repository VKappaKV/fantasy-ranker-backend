package domain

import (
	"context"
)

type PlayerRepository interface {
	UpsertByPUUID(ctx context.Context, p Player) (Player, error)
	ByPUUID(ctx context.Context, puuid string) (Player, error)
}
