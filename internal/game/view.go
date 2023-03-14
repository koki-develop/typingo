package game

func (g *Game) View() string {
	if g.showingResult {
		return "clear!!"
	}

	return g.words[g.currentWordIndex]
}
