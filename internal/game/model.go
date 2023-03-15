package game

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Game struct {
	// config
	words []string

	// state
	showingResult    bool
	currentWordIndex int
	currentCharIndex int
	windowWidth      int
	windowHeight     int
	duration         time.Duration

	// keymap
	keymap *KeyMap
}

type KeyMap struct {
	Cancel key.Binding
	Quit   key.Binding
}

type GameConfig struct {
	Words []string
}

var (
	_ tea.Model = (*Game)(nil)
)

func New(cfg *GameConfig) *Game {
	g := &Game{
		// config
		words: cfg.Words,

		// state
		showingResult:    false,
		currentWordIndex: 0,
		currentCharIndex: 0,
		duration:         0,

		// keymap
		keymap: &KeyMap{
			Cancel: key.NewBinding(
				key.WithKeys("ctrl+c", "esc"),
			),
			Quit: key.NewBinding(
				key.WithKeys("q"),
				key.WithHelp("q", "quit"),
			),
		},
	}

	g.keymap.Quit.SetEnabled(false)

	return g
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

func (g *Game) TypedChars() string {
	return string([]rune(g.CurrentWord())[:g.currentCharIndex])
}

func (g *Game) RemainChars() string {
	return string([]rune(g.CurrentWord())[g.currentCharIndex:])
}
