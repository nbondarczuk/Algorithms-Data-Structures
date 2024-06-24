package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryInsert(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("testing value %v->%v",
			test.args.tab, test.expected.tab),
			func(t *testing.T) {
				result := BinaryInsert(test.args.tab)
				assert.Equal(t, test.expected.tab, result)
			})
	}
}

func BenchmarkBinaryInsert(b *testing.B) {
	tab := []int{0, 1, 2, 3, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		BinaryInsert(tab)
	}
}
