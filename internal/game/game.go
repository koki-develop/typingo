package game

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type KeyMap struct {
	Cancel key.Binding
	Quit   key.Binding
}

/*
 * Model
 */

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

/*
 * Init
 */

func (g *Game) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		g.tick(),
	)
}

/*
 * View
 */

var (
	WordStyle = lipgloss.NewStyle().
			Align(lipgloss.Center, lipgloss.Center)
	TypedCharStyle = lipgloss.NewStyle().
			Faint(true)
	RemainCharStyle = lipgloss.NewStyle().
			Bold(true)
)

func (g *Game) View() string {
	view := ""

	view += g.duration.String()

	if g.showingResult {
		view += g.resultView()
	} else {
		view += g.wordView()
	}

	return view
}

func (g *Game) resultView() string {
	return "clear!!"
}

func (g *Game) wordView() string {
	typed := TypedCharStyle.Render(g.TypedChars())
	remain := RemainCharStyle.Render(g.RemainChars())
	word := lipgloss.JoinHorizontal(lipgloss.Center, typed, remain)

	return WordStyle.Width(g.windowWidth).Height(g.windowHeight).Render(word)
}

/*
 * Update
 */

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
