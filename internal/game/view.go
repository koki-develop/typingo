package game

import "github.com/charmbracelet/lipgloss"

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
	if g.showingResult {
		return "clear!!"
	}

	typed := TypedCharStyle.Render(g.TypedChars())
	remain := RemainCharStyle.Render(g.RemainChars())
	word := lipgloss.JoinHorizontal(lipgloss.Center, typed, remain)

	return WordStyle.Width(g.windowWidth).Height(g.windowHeight).Render(word)
}
