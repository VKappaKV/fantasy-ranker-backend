package domain

import "errors"

var (
	ErrInvalidRegion = errors.New("invalid region")
	ErrInvalidRiotID = errors.New("invalid riotId")
)
