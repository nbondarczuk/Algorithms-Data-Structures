package power

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFastPoweOf(t *testing.T) {
	for _, test := range tests {
		result := FastPowerOf(test.args.value, test.args.exp)
		if test.result.positive {
			assert.Equal(t, test.result.expected, result,
				fmt.Sprintf("%f ^ %d = %f but not %f",
					test.args.value, test.args.exp, test.result.expected, result))
		} else {
			assert.NotEqual(t, test.result.expected, result)
		}
	}
}

func BenchmarkFastPowerOfWith1MOfN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastPowerOf(1.0, 1000000)
	}
}
