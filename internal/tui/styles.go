package tui

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	App         lipgloss.Style
	Header      lipgloss.Style
	Panel       lipgloss.Style
	PanelTitle  lipgloss.Style
	Muted       lipgloss.Style
	Error       lipgloss.Style
	OK          lipgloss.Style
	Warn        lipgloss.Style
	Key         lipgloss.Style
	Value       lipgloss.Style
	Help        lipgloss.Style
	ResponseBox lipgloss.Style
}

func NewStyles() Styles {
	return Styles{
		App: lipgloss.NewStyle().Padding(1, 2),
		Header: lipgloss.NewStyle().
			Bold(true).
			Padding(0, 1),
		Panel: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(0, 1),
		PanelTitle: lipgloss.NewStyle().Bold(true),
		Muted:      lipgloss.NewStyle().Faint(true),
		Error:      lipgloss.NewStyle().Bold(true),
		OK:         lipgloss.NewStyle().Bold(true),
		Warn:       lipgloss.NewStyle().Bold(true),
		Key:        lipgloss.NewStyle().Bold(true),
		Value:      lipgloss.NewStyle(),
		Help:       lipgloss.NewStyle().Faint(true),
		ResponseBox: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(0, 1),
	}
}
