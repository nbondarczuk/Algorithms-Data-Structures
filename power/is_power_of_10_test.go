package power

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	CaseForPowerOf10 struct {
		value    float64
		expected bool
	}
)

var testsOfPowerOf10 = []CaseForPowerOf10{
	{
		1.0, true,
	},
	{
		0.1, true,
	},
	{
		10.0, true,
	},
	{
		3.0, false,
	},
	{
		0.3, false,
	},
}

func TestIsPowerOf10(t *testing.T) {
	for _, test := range testsOfPowerOf10 {
		result := IsPowerOf10(test.value)
		assert.Equal(t, test.expected, result,
			fmt.Sprintf("%v is a power of 10: %v", test.value, test.expected))
	}
}

func BenchmarkIsPowerOf10With1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPowerOf10(1.0)
	}
}
