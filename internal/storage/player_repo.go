package storage

import (
	"context"
	"errors"

	"github.com/VKappaKV/fantasy-ranker-backend/internal/domain/models"
)

type PlayerRepo struct {
	db *DB
}

func NewPlayerRepo(db *DB) *PlayerRepo {
	return &PlayerRepo{db: db}
}

func (r *PlayerRepo) UpsertByPUUID(ctx context.Context, p models.Player) (models.Player, error) {
	// Upsert: if puuid exists, update: gamename/tagline, region and updated_at. 
	// always returns the last row	
	const q = `
		INSERT INTO players (puuid, region, game_name, tag_line)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (puuid) DO UPDATE SET
			region = EXCLUDED.region,
			game_name = EXCLUDED.game_name,
			tag_line = EXCLUDED.tag_line,
			updated_at = NOW()
		RETURNING id, puuid, region, game_name, tag_line, created_at, updated_at;
		` 

	var out models.Player
	err := r.db.Pool.QueryRow(ctx, q, p.PUUID, string(p.Region), p.GameName, p.TagLine).Scan(&out.ID, &out.PUUID, &out.Region, &out.GameName, &out.TagLine, &out.CreatedAt, &out.UpdatedAt)
	if err != nil {
		return models.Player{}, err
	}

	if out.Region == "" {
		return models.Player{}, errors.New("invalid region stored in database")
	}
	return out, nil
}