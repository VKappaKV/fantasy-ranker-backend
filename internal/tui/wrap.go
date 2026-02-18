package tui

import (
	"strings"
	"unicode/utf8"
)

func wrapText(s string, width int) string {
	if width <= 0 {
		return s
	}

	lines := strings.Split(s, "\n")
	out := make([]string, 0, len(lines))

	for _, line := range lines {
		// keep empty lines
		if line == "" {
			out = append(out, "")
			continue
		}

		// simple rune-aware wrap (no fancy word-wrap; good enough for JSON)
		for utf8.RuneCountInString(line) > width {
			// cut by rune count
			runes := []rune(line)
			out = append(out, string(runes[:width]))
			line = string(runes[width:])
		}
		out = append(out, line)
	}

	return strings.Join(out, "\n")
}
