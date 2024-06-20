package search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinary(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("testing value %v,%v->%v@%v",
			test.args.tab, test.args.element, test.result.expected, test.result.index),
			func(t *testing.T) {
				index, found := Binary(test.args.tab, test.args.element)
				assert.Equal(t, test.result.expected, found)
				if found {
					assert.Equal(t, test.result.index, index)
				}
			})
	}
}
