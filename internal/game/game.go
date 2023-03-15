package game

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jedib0t/go-pretty/v6/text"
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
	count            int
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
		count:            3,
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

func (g *Game) wpm() float64 {
	chars := 0
	for _, w := range g.words {
		chars += utf8.RuneCountInString(w)
	}

	return g.endAt.Sub(g.startAt).Seconds() * 60
}

func (g *Game) running() bool {
	return g.count == 0
}

/*
 * Init
 */

func (g *Game) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		g.countdown(),
	)
}

/*
 * View
 */

var (
	centerStyle = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center)
)

func (g *Game) View() string {
	if !g.running() {
		return g.countdownView()
	}

	if g.showingResult {
		return g.resultView()
	} else {
		return g.wordView()
	}
}

func (g *Game) countdownView() string {
	return centerStyle.Width(g.windowWidth).Height(g.windowHeight).Render(strconv.Itoa(g.count))
}

func (g *Game) resultView() string {
	view := ""

	view += lipgloss.NewStyle().Bold(true).Render("Result") + "\n\n"

	rows := []string{
		fmt.Sprintf("Record: %s", g.endAt.Sub(g.startAt).Truncate(time.Millisecond).String()),
		fmt.Sprintf("Miss:   %d", g.miss),
		fmt.Sprintf("WPM:    %d", int(g.wpm())),
	}
	maxlen := text.LongestLineLen(strings.Join(rows, "\n"))
	for _, row := range rows {
		view += text.Pad(row, maxlen, ' ') + "\n"
	}
	view += "\n"

	view += "[r] retry" + "\n"
	view += "[q] quit " + "\n"

	return centerStyle.Height(g.windowHeight).Width(g.windowWidth).Render(
		view,
	)
}

func (g *Game) wordView() string {
	typed := lipgloss.NewStyle().Faint(true).Render(g.typedChars())
	remain := lipgloss.NewStyle().Bold(true).Render(g.remainChars())

	return centerStyle.Width(g.windowWidth).Height(g.windowHeight).Render(
		lipgloss.JoinHorizontal(lipgloss.Center, typed, remain),
	)
}

/*
 * Update
 */

type countdownMsg struct{}

func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, g.keymap.Cancel):
			return g, tea.Quit
		case key.Matches(msg, g.keymap.Quit):
			return g, tea.Quit
		case g.running() && !g.showingResult:
			g.pressKey(msg)
		}
	case countdownMsg:
		g.count--
		if g.running() {
			g.startAt = time.Now()
		} else {
			return g, g.countdown()
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

func (g *Game) countdown() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return countdownMsg{}
	})
}
