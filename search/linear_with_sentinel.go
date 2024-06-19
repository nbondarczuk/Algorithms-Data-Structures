package search

// Cost O(N/2) where N = len(tab)
func LinearWithSentinel(tab []int, e int) (int, bool) {
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
