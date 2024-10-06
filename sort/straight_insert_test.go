package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStraightInsert(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("testing value %v->%v",
			test.args.tab, test.expected.tab),
			func(t *testing.T) {
				tab := make([]int, len(test.args.tab), len(test.args.tab))
				copy(tab, test.args.tab)
				result := StraightInsert(tab)
				assert.Equal(t, test.expected.tab, result)
			})
	}
}

func BenchmarkStraightInsert(b *testing.B) {
	tab := []int{0, 1, 2, 3, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		StraightInsert(tab)
	}
}
