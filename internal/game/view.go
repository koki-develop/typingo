package game

func (g *Game) View() string {
	return g.words[g.currentWordIndex]
}
