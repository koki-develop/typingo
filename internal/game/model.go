package game

import tea "github.com/charmbracelet/bubbletea"

type Game struct {
	words []string
}

type GameConfig struct {
	Words []string
}

var (
	_ tea.Model = (*Game)(nil)
)

func New(cfg *GameConfig) *Game {
	return &Game{
		words: cfg.Words,
	}
}

func Run(g *Game) error {
	p := tea.NewProgram(g)
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}
