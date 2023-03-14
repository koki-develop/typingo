package game

import tea "github.com/charmbracelet/bubbletea"

type Game struct{}

var (
	_ tea.Model = (*Game)(nil)
)

func New() *Game {
	return &Game{}
}

func Run(g *Game) error {
	p := tea.NewProgram(g)
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}
