package text

// Thr cost is O(N*M)
func SearchSimple(s, p string) int {
	var i int
	n := len(s)
	m := len(p)
	for i = 0; i <= n-m; i++ {
		if s[i:i+m] == p {
			return i
		}
	}
	return -1
}
