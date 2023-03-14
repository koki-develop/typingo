package game

import (
	"github.com/charmbracelet/lipgloss"
)

// styles
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
