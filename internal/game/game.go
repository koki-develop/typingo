package game

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keymap struct {
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
	miss             int
	currentWordIndex int
	currentCharIndex int
	windowWidth      int
	windowHeight     int
	startAt          time.Time
	endAt            time.Time

	// keymap
	keymap *keymap
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
		miss:             0,
		currentWordIndex: 0,
		currentCharIndex: 0,

		// keymap
		keymap: &keymap{
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

func (g *Game) currentWord() string {
	return g.words[g.currentWordIndex]
}

func (g *Game) currentChar() string {
	return string([]rune(g.currentWord())[g.currentCharIndex])
}

func (g *Game) typedChars() string {
	return string([]rune(g.currentWord())[:g.currentCharIndex])
}

func (g *Game) remainChars() string {
	return string([]rune(g.currentWord())[g.currentCharIndex:])
}

/*
 * Init
 */

func (g *Game) Init() tea.Cmd {
	g.startAt = time.Now()

	return tea.Batch(
		tea.EnterAltScreen,
	)
}

/*
 * View
 */

var (
	wordStyle = lipgloss.NewStyle().
			Align(lipgloss.Center, lipgloss.Center)
	typedCharStyle = lipgloss.NewStyle().
			Faint(true)
	remainCharStyle = lipgloss.NewStyle().
			Bold(true)

	resultStyle = lipgloss.NewStyle().
			Align(lipgloss.Center, lipgloss.Center)
	resultHeadingStyle = lipgloss.NewStyle().
				Bold(true).
				Padding(1)
	resultDurationStyle = lipgloss.NewStyle().MarginBottom(1)
	resultHelpStyle     = lipgloss.NewStyle()
)

func (g *Game) View() string {
	view := ""

	if g.showingResult {
		view += g.resultView()
	} else {
		view += g.wordView()
	}

	return view
}

func (g *Game) resultView() string {
	heading := resultHeadingStyle.Render("Result")
	duration := resultDurationStyle.Render(fmt.Sprintf(
		"Record: %s\nMiss: %d",
		g.endAt.Sub(g.startAt).Truncate(time.Millisecond).String(),
		g.miss,
	))
	help := resultHelpStyle.Render("Press q to quit")

	return resultStyle.Width(g.windowWidth).Height(g.windowHeight).Render(
		lipgloss.JoinVertical(lipgloss.Center, heading, duration, help),
	)
}

func (g *Game) wordView() string {
	typed := typedCharStyle.Render(g.typedChars())
	remain := remainCharStyle.Render(g.remainChars())
	word := lipgloss.JoinHorizontal(lipgloss.Center, typed, remain)

	return wordStyle.Width(g.windowWidth).Height(g.windowHeight).Render(word)
}

/*
 * Update
 */

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
	}

	return g, nil
}

func (g *Game) pressKey(msg tea.KeyMsg) {
	if msg.String() == g.currentChar() {
		g.currentCharIndex++

		if g.currentCharIndex == len(g.currentWord()) {
			g.currentCharIndex = 0
			g.currentWordIndex++

			if g.currentWordIndex == len(g.words) {
				g.endAt = time.Now()
				g.showingResult = true
				g.keymap.Quit.SetEnabled(true)
			}
		}
	} else {
		g.miss++
	}
}
