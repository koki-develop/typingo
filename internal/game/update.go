package game

import tea "github.com/charmbracelet/bubbletea"

func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return g, tea.Quit
		}
	}

	return g, nil
}
