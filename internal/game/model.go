package game

import tea "github.com/charmbracelet/bubbletea"

type Game struct {
	// config
	words []string

	// state
	showingResult    bool
	currentWordIndex int
	currentCharIndex int
	windowWidth      int
	windowHeight     int
}

type GameConfig struct {
	Words []string
}

var (
	_ tea.Model = (*Game)(nil)
)

func New(cfg *GameConfig) *Game {
	return &Game{
		words:            cfg.Words,
		showingResult:    false,
		currentWordIndex: 0,
		currentCharIndex: 0,
	}
}

func Run(g *Game) error {
	p := tea.NewProgram(g)
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

func (g *Game) CurrentWord() string {
	return g.words[g.currentWordIndex]
}

func (g *Game) CurrentChar() string {
	return string([]rune(g.CurrentWord())[g.currentCharIndex])
}
