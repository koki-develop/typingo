package texts

import (
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
)

var fakeFns = []func() string{
	gofakeit.Phrase,
}

func Random(l int) []string {
	texts := make([]string, l)
	for i := 0; i < l; i++ {
		texts[i] = fakeFns[rand.Intn(len(fakeFns))]()
	}
	return texts
}
