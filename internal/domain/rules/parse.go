package rules

import (
	"fmt"
	"strings"

	"github.com/VKappaKV/fantasy-ranker-backend/internal/domain/models"
)

func ParseRiotID(s string) (models.RiotID, error) {
	parts := strings.Split(strings.TrimSpace(s), "#")
	if len(parts) != 2 {
		return models.RiotID{}, fmt.Errorf("%w", ErrInvalidRiotID)
	}
	gn := strings.TrimSpace(parts[0])
	tl := strings.TrimSpace(parts[1])
	if gn == "" || tl == "" {
		return models.RiotID{}, fmt.Errorf("%w", ErrInvalidRiotID)
	}
	return models.RiotID{GameName: gn, TagLine: tl}, nil
}

func ParseRegion(s string) (models.Region, error) {
	v := models.Region(strings.ToLower(strings.TrimSpace(s)))
	switch v {
	case models.RegionEurope, models.RegionAmericas, models.RegionAsia:
		return v, nil
	default:
		return "", fmt.Errorf("%w: %q", ErrInvalidRegion, s)
	}
}