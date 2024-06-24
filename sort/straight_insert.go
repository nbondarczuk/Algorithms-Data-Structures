package sort

import "golang.org/x/exp/constraints"

func StraightInsert[T constraints.Ordered](tab []T) []T {
	var i, j int
	n := len(tab)
	for i = 1; i <= n-1; i++ {
		x := tab[i]
		j = i
		for j > 0 && x < tab[j-1] {
			tab[j] = tab[j-1]
			j--
		}
		tab[j] = x
	}
	return tab
}
