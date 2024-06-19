package search

// Cost O(N/2) where N = len(tab)
func Linear(tab []int, e int) (int, bool) {
	var i int
	n := len(tab)
	for {
		if i >= n || tab[i] == e {
			break
		}
		i++
	}
	return i, i < n
}
