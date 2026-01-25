package domain

import (
	"fmt"
	"strings"
)

type RiotID struct {
	GameName string
	TagLine  string
}

func ParseRiotID(s string) (RiotID, error) {
	parts := strings.Split(strings.TrimSpace(s), "#")
	if len(parts) != 2 {
		return RiotID{}, fmt.Errorf("%w", ErrInvalidRiotID)
	}
	gn := strings.TrimSpace(parts[0])
	tl := strings.TrimSpace(parts[1])
	if gn == "" || tl == "" {
		return RiotID{}, fmt.Errorf("%w", ErrInvalidRiotID)
	}
	return RiotID{GameName: gn, TagLine: tl}, nil
}
