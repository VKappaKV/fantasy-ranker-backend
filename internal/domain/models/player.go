package models

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
)

type Player struct {
	ID        uuid.UUID
	PUUID     string
	Region    Region
	GameName  string
	TagLine   string
	CreatedAt time.Time
	UpdatedAt time.Time
}