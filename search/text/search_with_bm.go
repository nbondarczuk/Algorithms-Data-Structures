package text

// Text search Boyer-Moore algorithm
func SearchWithBM(s, p string) int {
	var i, j, k, r int
	n := len(s)
	m := len(p)
	d := [128]int{}
	for i = 0; i <= 127; i++ {
		d[i] = m
	}
	for j = 0; j <= m-2; j++ {
		d[uint8(p[j])] = m - j - 1
	}
	i = m
	for {
		j = m
		k = i
		for {
			k--
			j--
			if j < 0 || p[j] != s[k] {
				break
			}
		}
		i = i + d[uint8(s[i-1])]
		if j < 0 || i > n {
			break
		}
	}
	if j < 0 {
		r = k + 1
	} else {
		r = -1
	}
	return r
}
