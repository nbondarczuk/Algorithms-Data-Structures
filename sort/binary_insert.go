package sort

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func bsearch[T constraints.Ordered](tab []T, L, R int, x T) int {
	var m int
	for L <= R {
		m = L + (R-L)/2
		if x < tab[m] {
			R = m - 1
		} else {
			L = m + 1
		}
	}
	return L
}

func BinaryInsert[T constraints.Ordered](tab []T) []T {
	var i, j, n, p int
	n = len(tab)
	for i = 1; i < n; i++ {
		x := tab[i]
		p = bsearch(tab, 0, i-1, x)
		for j = i - 1; j >= p; j-- {
			tab[j+1] = tab[j]
		}
		tab[p] = x
		if Debug {
			fmt.Printf("Tab: %+v\n", tab)
		}
	}
	return tab
}
