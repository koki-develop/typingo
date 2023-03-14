package game

import tea "github.com/charmbracelet/bubbletea"

func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return g, tea.Quit
		default:
			key := msg.String()
			if g.showingResult {
				if key == "q" {
					return g, tea.Quit
				}
			} else {
				g.pressKey(key)
			}
		}
	}

	return g, nil
}

func (g *Game) pressKey(key string) {
	currentWord := g.words[g.currentWordIndex]
	currentChar := string([]rune(currentWord)[g.currentCharIndex])

	if currentChar == key {
		g.currentCharIndex++

		if g.currentCharIndex == len(currentWord) {
			g.currentCharIndex = 0
			g.currentWordIndex++

			if g.currentWordIndex == len(g.words) {
				g.showingResult = true
			}
		}
	}
}
