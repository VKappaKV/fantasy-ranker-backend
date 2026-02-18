package main

import (
	"log"
	"os"

	"github.com/VKappaKV/fantasy-ranker-backend/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	baseURL := os.Getenv("FANTASY_BACKEND_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	p := tea.NewProgram(tui.NewModel(baseURL), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
