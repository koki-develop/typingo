package game

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (g *Game) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		g.tick(),
	)
}
