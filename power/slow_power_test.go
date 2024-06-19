package power

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlowPoweOf(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("testing value %v^%v->%v", test.args.value, test.args.exp, test.result.expected), func(t *testing.T) {
			result := SlowPowerOf(test.args.value, test.args.exp)
			if test.result.positive {
				assert.Equal(t, test.result.expected, result,
					fmt.Sprintf("%f ^ %d = %f but not %f",
						test.args.value, test.args.exp, test.result.expected, result))
			} else {
				assert.NotEqual(t, test.result.expected, result)
			}
		})
	}
}

func BenchmarkSlowPowerOfWith1MOfN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SlowPowerOf(1.0, 1000000)
	}
}
