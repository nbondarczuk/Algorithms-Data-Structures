package sort

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Count[T constraints.Ordered](tab []T) []T {
	count := make([]int, len(tab), len(tab))
	n := len(tab)
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if j != i && tab[j] < tab[i] {
				count[i]++
			}
		}
		if Debug {
			fmt.Printf("Count: %+v\n", count)
		}
	}
	rval := make([]T, len(tab), len(tab))
	for i := 0; i < n; i++ {
		rval[i] = tab[count[i]]
	}
	return rval
}
