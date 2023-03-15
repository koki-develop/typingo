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
	"github.com/koki-develop/typingo/internal/texts"
)

type keymap struct {
	Start  key.Binding
	Retry  key.Binding
	Cancel key.Binding
	Quit   key.Binding
}

/*
 * Model
 */

type Game struct {
	// config
	numTexts int

	// state
	texts            []string
	start            bool
	count            int
	miss             int
	currentTextIndex int
	currentCharIndex int
	startAt          time.Time
	endAt            time.Time

	windowWidth  int
	windowHeight int

	// keymap
	keymap *keymap
}

type GameConfig struct {
	NumTexts int
}

var (
	_ tea.Model = (*Game)(nil)
)

func New(cfg *GameConfig) *Game {
	g := &Game{
		// config
		numTexts: cfg.NumTexts,

		// keymap
		keymap: &keymap{
			Start: key.NewBinding(
				key.WithKeys(" "),
			),
			Retry: key.NewBinding(
				key.WithKeys("r"),
			),
			Cancel: key.NewBinding(
				key.WithKeys("ctrl+c", "esc"),
			),
			Quit: key.NewBinding(
				key.WithKeys("q"),
				key.WithHelp("q", "quit"),
			),
		},
	}

	g.reset()
	return g
}

func Run(g *Game) error {
	p := tea.NewProgram(g)
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

func (g *Game) reset() {
	// state
	g.texts = texts.Random(g.numTexts)
	g.start = false
	g.count = 3
	g.miss = 0
	g.currentTextIndex = 0
	g.currentCharIndex = 0

	// keymap
	g.keymap.Start.SetEnabled(true)
	g.keymap.Retry.SetEnabled(false)
	g.keymap.Quit.SetEnabled(false)
}

func (g *Game) currentText() string {
	return g.texts[g.currentTextIndex]
}

func (g *Game) currentChar() string {
	return string([]rune(g.currentText())[g.currentCharIndex])
}

func (g *Game) typedChars() string {
	return string([]rune(g.currentText())[:g.currentCharIndex])
}

func (g *Game) remainChars() string {
	return string([]rune(g.currentText())[g.currentCharIndex+1:])
}

func (g *Game) wpm() float64 {
	chars := 0
	for _, w := range g.texts {
		chars += utf8.RuneCountInString(w)
	}

	return float64(chars) / g.endAt.Sub(g.startAt).Seconds() * 60
}

func (g *Game) showingResult() bool {
	return g.currentTextIndex == len(g.texts)
}

func (g *Game) running() bool {
	return g.count == 0 && !g.showingResult()
}

/*
 * Init
 */

func (g *Game) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
	)
}

/*
 * View
 */

var (
	mainColor = lipgloss.Color("#00ADD8")
)

func newCenterStyle() lipgloss.Style {
	return lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center)
}

func pad(s string) string {
	maxlen := text.LongestLineLen(s)
	rows := strings.Split(s, "\n")
	rslt := make([]string, len(rows))
	for i := 0; i < len(rows); i++ {
		rslt[i] = text.Pad(rows[i], maxlen, ' ')
	}
	return strings.Join(rslt, "\n")
}

func (g *Game) View() string {
	switch {
	case !g.start:
		return g.startView()
	case g.running():
		return g.textView()
	case g.showingResult():
		return g.resultView()
	default:
		return g.countdownView()
	}
}

func (g *Game) startView() string {
	logo := pad(` ________                      __
/        |                    /  |
$$$$$$$$/  __    __   ______  $$/  _______    ______    ______
	 $$ |   /  |  /  | /      \ /  |/       \  /      \  /      \
	 $$ |   $$ |  $$ |/$$$$$$  |$$ |$$$$$$$  |/$$$$$$  |/$$$$$$  |
	 $$ |   $$ |  $$ |$$ |  $$ |$$ |$$ |  $$ |$$ |  $$ |$$ |  $$ |
	 $$ |   $$ \__$$ |$$ |__$$ |$$ |$$ |  $$ |$$ \__$$ |$$ \__$$ |
	 $$ |   $$    $$ |$$    $$/ $$ |$$ |  $$ |$$    $$ |$$    $$/
	 $$/     $$$$$$$ |$$$$$$$/  $$/ $$/   $$/  $$$$$$$ | $$$$$$/
					/  \__$$ |$$ |                    /  \__$$ |
					$$    $$/ $$ |                    $$    $$/
					 $$$$$$/  $$/                      $$$$$$/
`)

	return newCenterStyle().Width(g.windowWidth).Height(g.windowHeight).Render(
		lipgloss.NewStyle().Bold(true).Foreground(mainColor).Render(logo),
		"\n",
		lipgloss.NewStyle().Render("press space to start"),
	)
}

func (g *Game) countdownView() string {
	return newCenterStyle().Width(g.windowWidth).Height(g.windowHeight).Bold(true).Render(strconv.Itoa(g.count))
}

func (g *Game) resultView() string {
	view := ""

	view += lipgloss.NewStyle().Foreground(mainColor).Bold(true).Render("Result") + "\n\n"

	view += lipgloss.NewStyle().Bold(true).Render(pad(
		fmt.Sprintf("Record: %s", g.endAt.Sub(g.startAt).Truncate(time.Millisecond).String()) + "\n" +
			fmt.Sprintf("Miss:   %d", g.miss) + "\n" +
			fmt.Sprintf("WPM:    %d", int(g.wpm())),
	))
	view += "\n\n"

	view += "[r] retry" + "\n"
	view += "[q] quit " + "\n"

	return newCenterStyle().Height(g.windowHeight).Width(g.windowWidth).Render(
		view,
	)
}

func (g *Game) textView() string {
	view := ""

	typed := lipgloss.NewStyle().Faint(true).Render(g.typedChars())
	char := lipgloss.NewStyle().Bold(true).Underline(true).Render(g.currentChar())
	remain := lipgloss.NewStyle().Bold(true).Render(g.remainChars())
	view += lipgloss.JoinHorizontal(lipgloss.Center, typed, char, remain)
	view += "\n"

	view += fmt.Sprintf("(%d/%d)", g.currentTextIndex+1, len(g.texts))

	return newCenterStyle().Width(g.windowWidth).Height(g.windowHeight).Render(
		view,
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
		case key.Matches(msg, g.keymap.Start):
			g.start = true
			g.keymap.Start.SetEnabled(false)
			return g, g.countdown()
		case g.running():
			g.pressKey(msg)
		case g.showingResult() && key.Matches(msg, g.keymap.Retry):
			g.reset()
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

		if g.currentCharIndex == len(g.currentText()) {
			g.currentCharIndex = 0
			g.currentTextIndex++

			if g.currentTextIndex == len(g.texts) {
				g.endAt = time.Now()
				g.keymap.Retry.SetEnabled(true)
				g.keymap.Quit.SetEnabled(true)
			}
		}
	} else {
		g.miss++
		fmt.Print("\a")
	}
}

func (g *Game) countdown() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return countdownMsg{}
	})
}
