package search

// Cost is O(log N)
func Binary[T ~byte | ~rune | ~int | ~uint | ~float32 | ~float64 | ~string](tab []T, e T) (int, bool) {
	var left, right, middle int
	right = len(tab) - 1
	for {
		if left >= right {
			break
		}
		middle = (left + right) / 2
		if tab[middle] < e {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return right, tab[right] == e
}
