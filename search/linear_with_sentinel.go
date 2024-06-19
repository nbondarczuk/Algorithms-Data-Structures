package search

// Cost O(N/2) where N = len(tab)
func LinearWithSentinel[T comparable](tab []T, e T) (int, bool) {
	stab := append(tab, e)
	var i int
	n := len(tab)
	for {
		if stab[i] == e {
			break
		}
		i++
	}
	return i, i < n
}
