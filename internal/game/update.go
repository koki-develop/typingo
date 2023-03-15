package game

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg struct{}

func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, g.keymap.Cancel):
			return g, tea.Quit
		case key.Matches(msg, g.keymap.Quit):
			return g, tea.Quit
		case !g.showingResult:
			g.pressKey(msg)
		}
	case tea.WindowSizeMsg:
		g.windowWidth, g.windowHeight = msg.Width, msg.Height
	case TickMsg:
		if g.showingResult {
			break
		}
		g.duration += time.Millisecond
		return g, g.tick()
	}

	return g, nil
}

func (g *Game) pressKey(msg tea.KeyMsg) {
	if msg.String() == g.CurrentChar() {
		g.currentCharIndex++

		if g.currentCharIndex == len(g.CurrentWord()) {
			g.currentCharIndex = 0
			g.currentWordIndex++

			if g.currentWordIndex == len(g.words) {
				g.showingResult = true
				g.keymap.Quit.SetEnabled(true)
			}
		}
	}
}

func (g *Game) tick() tea.Cmd {
	return tea.Tick(time.Millisecond, func(t time.Time) tea.Msg {
		return TickMsg{}
	})
}
