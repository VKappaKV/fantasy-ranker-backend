package tui

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type endpointItem Endpoint

func (e endpointItem) Title() string       { return e.Name }
func (e endpointItem) Description() string { return e.Help }
func (e endpointItem) FilterValue() string { return e.Name + " " + e.Path }

type callResultMsg struct {
	resp *APIResponse
	err  error
}

type Model struct {
	baseURL string
	client  *Client
	styles  Styles

	width  int
	height int

	// left side
	endpoints []Endpoint
	list      list.Model

	// param builder
	paramInputs []textinput.Model
	paramSpecs  []ParamSpec
	paramHint   string

	// right side
	vp        viewport.Model
	lastResp  *APIResponse
	lastError error

	// ui state
	focusIdx int // which param input focused
}

func NewModel(baseURL string) Model {
	s := NewStyles()
	c := NewClient(baseURL)

	endpoints := []Endpoint{
		{
			Name:        "Health",
			Method:      "GET",
			Path:        "/health",
			Help:        "Backend health check",
		},
		{
			Name:        "Version",
			Method:      "GET",
			Path:        "/version",
			Help:        "Backend version info",
		},
		{
			Name:        "Account",
			Method:      "GET",
			Path:        "/account",
			Help:        "Lookup account (requires params)",
			Params: []ParamSpec{
				{Key: "gameName", Required: true, Hint: "Riot ID game name"},
				{Key: "tagLine", Required: true, Hint: "Riot ID tagline (e.g. EUW)"},
				// add others if your backend expects them
			},
		},
		{
			Name:        "Matches",
			Method:      "GET",
			Path:        "/matches",
			Help:        "Fetch matches (requires params)",
			Params: []ParamSpec{
				{Key: "puuid", Required: true, Hint: "Player PUUID"},
				{Key: "count", Required: false, Hint: "How many matches"},
				// add start, queue, type, etc. as needed
			},
		},
	}

	items := make([]list.Item, 0, len(endpoints))
	for _, e := range endpoints {
		items = append(items, endpointItem(e))
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Endpoints"
	l.SetShowHelp(false)

	m := Model{
		baseURL:   baseURL,
		client:    c,
		styles:    s,
		endpoints: endpoints,
		list:      l,
		vp:        viewport.New(0, 0),
	}

	m.loadSelectedEndpointParams()
	return m
}

func (m *Model) loadSelectedEndpointParams() {
	it, ok := m.list.SelectedItem().(endpointItem)
	if !ok {
		m.paramSpecs = nil
		m.paramInputs = nil
		return
	}
	specs := []ParamSpec(it.Params)
	m.paramSpecs = specs
	m.paramInputs = make([]textinput.Model, len(specs))

	for i, p := range specs {
		ti := textinput.New()
		ti.Prompt = ""
		ti.Placeholder = p.Key
		ti.CharLimit = 256
		ti.Width = 28
		if p.Required {
			ti.Placeholder = p.Key + " *"
		}
		m.paramInputs[i] = ti
	}
	m.focusIdx = 0
	if len(m.paramInputs) > 0 {
		m.paramInputs[0].Focus()
		m.paramHint = specs[0].Hint
	} else {
		m.paramHint = ""
	}
}

func (m *Model) refreshViewportContent() {
	content := m.renderResponseContent()

	// Wrap to viewport width so long JSON lines don't explode layout.
	// (If you use padding inside response box, subtract 0/1 accordingly)
	wrapped := wrapText(content, m.vp.Width)

	m.vp.SetContent(wrapped)

	// If you want "stick to top on new response":
	// m.vp.GotoTop()
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
	m.width, m.height = msg.Width, msg.Height

	// overall outer padding from styles.App (Padding(1,2)) => 2 lines vertical, 4 cols horizontal
	outerV := 2
	outerH := 4

	usableW := m.width - outerH
	usableH := m.height - outerV

	// header uses 1 line + blank line (you render "\n\n" after header)
	headerH := 2

	bodyH := usableH - headerH
	if bodyH < 5 {
		bodyH = 5
	}

	leftW := max(40, usableW/3)
	gutter := 3
	rightW := usableW - leftW - gutter
	if rightW < 30 {
		rightW = 30
		leftW = usableW - rightW - gutter
	}

	// Right panel has Params + Response box
	// Estimate params box height:
	// border(2) + title line(1) + inputs(len + hint + enter line)
	paramsInner := 1 + len(m.paramSpecs) + 2 // rough: title + inputs + hint/enter
	paramsH := paramsInner + 2               // borders
	if paramsH < 6 {
		paramsH = 6
	}
	// Keep params from eating everything
	if paramsH > bodyH/2 {
		paramsH = bodyH / 2
	}

	// Response box gets the rest
	responseBoxH := bodyH - paramsH - 1 // minus the newline you add between params and response
	if responseBoxH < 6 {
		responseBoxH = 6
	}

	// List height: left panel border(2) + title/foot (2-ish)
	m.list.SetSize(leftW-2, bodyH-4)

	// Viewport lives INSIDE response box:
	// response box has border(2) + title line(1) + help line(1) + blank line spacing maybe
	vpH := responseBoxH - 2 - 1 - 1
	if vpH < 1 {
		vpH = 1
	}
	// and width inside border/padding
	vpW := rightW - 2 - 2 // border(2) + padding(0,1) => 2 columns
	if vpW < 10 {
		vpW = 10
	}

	m.vp.Width = vpW
	m.vp.Height = vpH

	m.refreshViewportContent()
	return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "tab":
			// focus next param input (if any)
			if len(m.paramInputs) > 0 {
				m.paramInputs[m.focusIdx].Blur()
				m.focusIdx = (m.focusIdx + 1) % len(m.paramInputs)
				m.paramInputs[m.focusIdx].Focus()
				m.paramHint = m.paramSpecs[m.focusIdx].Hint
			}
			return m, nil

		case "shift+tab":
			if len(m.paramInputs) > 0 {
				m.paramInputs[m.focusIdx].Blur()
				m.focusIdx = (m.focusIdx - 1 + len(m.paramInputs)) % len(m.paramInputs)
				m.paramInputs[m.focusIdx].Focus()
				m.paramHint = m.paramSpecs[m.focusIdx].Hint
			}
			return m, nil

		case "enter":
			// If focus is in list and user hits enter, we call endpoint.
			// If focus is in params, also call endpoint (quick workflow).
			return m, m.callSelectedEndpoint()

		case "r":
			// reset param values
			for i := range m.paramInputs {
				m.paramInputs[i].SetValue("")
			}
			m.lastResp = nil
			m.lastError = nil
			m.vp.SetContent(m.renderResponseContent())
			return m, nil
		}

	case callResultMsg:
		m.lastResp = msg.resp
		m.lastError = msg.err
		m.refreshViewportContent()
		return m, nil
	}

	// Update list
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	// If selection changed, reload params
	// (list delegate doesn't give an explicit "selection changed" msg,
	// so detect via keypresses that change selection by comparing selected endpoint)
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
		case "up", "down", "k", "j", "pgup", "pgdown", "home", "end":
			m.loadSelectedEndpointParams()
		}
	}

	// Update textinputs
	for i := range m.paramInputs {
		m.paramInputs[i], cmd = m.paramInputs[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	// Update viewport (scrolling)
	m.vp, cmd = m.vp.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) callSelectedEndpoint() tea.Cmd {
	it, ok := m.list.SelectedItem().(endpointItem)
	if !ok {
		return nil
	}

	// validate required params
	query := map[string]string{}
	for i, spec := range m.paramSpecs {
		val := strings.TrimSpace(m.paramInputs[i].Value())
		if spec.Required && val == "" {
			return func() tea.Msg {
				return callResultMsg{resp: nil, err: fmt.Errorf("missing required param: %s", spec.Key)}
			}
		}
		if val != "" {
			query[spec.Key] = val
		}
	}

	ep := Endpoint(it)
	return func() tea.Msg {
		resp, err := m.client.Do(context.Background(), ep.Method, ep.Path, query, nil)
		return callResultMsg{resp: resp, err: err}
	}
}

func (m Model) View() string {
	s := m.styles

	header := s.Header.Render(
		fmt.Sprintf("Fantasy Ranker TUI  •  %s", s.Muted.Render(m.baseURL)),
	)

	leftW := max(40, m.width/3)
	rightW := m.width - leftW - 3

	left := lipgloss.NewStyle().Width(leftW).Render(
		s.Panel.Render(
			s.PanelTitle.Render("Endpoints")+"\n"+
				m.list.View()+"\n"+
				s.Muted.Render("↑/↓ select • Enter call • r reset"),
		),
	)

	params := s.Panel.Render(
		s.PanelTitle.Render("Params")+" "+s.Muted.Render("(Tab to move)")+"\n"+
			m.renderParams(),
	)

	right := lipgloss.NewStyle().Width(rightW).Render(
		params + "\n" +
			s.ResponseBox.Render(
				s.PanelTitle.Render("Response")+"\n"+
					m.vp.View()+"\n"+
					s.Help.Render("Scroll: ↑/↓ • q quit"),
			),
	)

	body := lipgloss.JoinHorizontal(lipgloss.Top, left, "   ", right)

	return s.App.Render(header + "\n\n" + body)
}

func (m Model) renderParams() string {
	if len(m.paramSpecs) == 0 {
		return m.styles.Muted.Render("No parameters for this endpoint.\n\nEnter to call.")
	}

	lines := make([]string, 0, len(m.paramSpecs)+2)
	for i, spec := range m.paramSpecs {
		label := spec.Key
		if spec.Required {
			label += " *"
		}
		label = m.styles.Key.Render(label)

		input := m.paramInputs[i].View()
		lines = append(lines, fmt.Sprintf("%-14s %s", label, input))
	}
	if m.paramHint != "" {
		lines = append(lines, "")
		lines = append(lines, m.styles.Muted.Render("Hint: "+m.paramHint))
	}
	lines = append(lines, m.styles.Muted.Render("Enter to call endpoint."))

	return strings.Join(lines, "\n")
}

func (m Model) renderResponseContent() string {
	if m.lastError != nil {
		return m.styles.Error.Render("Error: ") + m.lastError.Error()
	}
	if m.lastResp == nil {
		return m.styles.Muted.Render("No response yet. Select an endpoint, fill params (if any), and press Enter.")
	}

	summary := m.lastResp.SummaryLine()
	statusLine := summary
	if m.lastResp.StatusCode >= 200 && m.lastResp.StatusCode < 300 {
		statusLine = m.styles.OK.Render(statusLine)
	} else if m.lastResp.StatusCode >= 400 {
		statusLine = m.styles.Error.Render(statusLine)
	} else {
		statusLine = m.styles.Warn.Render(statusLine)
	}

	body := m.lastResp.BodyPretty
	if strings.TrimSpace(body) == "" {
		body = m.styles.Muted.Render("(empty body)")
	}

	return statusLine + "\n\n" + body
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
