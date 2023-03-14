package game

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg struct{}

func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return g, tea.Quit
		default:
			key := msg.String()
			if g.showingResult {
				if key == "q" {
					return g, tea.Quit
				}
			} else {
				g.pressKey(key)
			}
		}
	case tea.WindowSizeMsg:
		g.windowWidth, g.windowHeight = msg.Width, msg.Height
	case TickMsg:
		if g.showingResult {
			break
		}
		return g, g.tick()
	}

	return g, nil
}

func (g *Game) pressKey(key string) {
	if key == g.CurrentChar() {
		g.currentCharIndex++

		if g.currentCharIndex == len(g.CurrentWord()) {
			g.currentCharIndex = 0
			g.currentWordIndex++

			if g.currentWordIndex == len(g.words) {
				g.showingResult = true
			}
		}
	}
}

func (g *Game) tick() tea.Cmd {
	return tea.Tick(time.Millisecond, func(t time.Time) tea.Msg {
		g.duration += time.Millisecond
		return TickMsg{}
	})
}
