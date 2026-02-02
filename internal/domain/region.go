package domain

import (
	"fmt"
	"strings"
)

type Region string

const (
	RegionEurope   Region = "europe"
	RegionAmericas Region = "americas"
	RegionAsia     Region = "asia"
)

func ParseRegion(s string) (Region, error) {
	v := Region(strings.ToLower(strings.TrimSpace(s)))
	switch v {
	case RegionEurope, RegionAmericas, RegionAsia:
		return v, nil
	default:
		return "", fmt.Errorf("%w: %q", ErrInvalidRegion, s)
	}
}