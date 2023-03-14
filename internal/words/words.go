package words

import "github.com/brianvoe/gofakeit/v6"

func Random(l int) []string {
	words := make([]string, l)
	for i := 0; i < l; i++ {
		words[i] = gofakeit.Noun()
	}
	return words
}
