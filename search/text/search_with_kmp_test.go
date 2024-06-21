package text

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchWithKMP(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("testing value %v,%v->%v",
			test.args.text, test.args.word, test.result.index),
			func(t *testing.T) {
				index := SearchWithKMP(test.args.text, test.args.word)
				assert.Equal(t, test.result.index, index)
			})
	}
}

func BenchmarkSearchWithKMP(b *testing.B) {
	text := "alfa beta gamma delta iota kappa"
	word := "alfa"
	for i := 0; i < b.N; i++ {
		SearchWithKMP(text, word)
	}
}
