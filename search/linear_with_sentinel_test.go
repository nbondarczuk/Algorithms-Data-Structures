package search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearWithSentinel(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("testing value %v,%v->%v@%v",
			test.args.tab, test.args.element, test.result.expected, test.result.index),
			func(t *testing.T) {
				index, found := LinearWithSentinel(test.args.tab, test.args.element)
				assert.Equal(t, test.result.expected, found)
				assert.Equal(t, test.result.index, index)
			})
	}
}
