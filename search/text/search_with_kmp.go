package text

const Mmax = 1024

// Text search Knuth-Morris-Pratt algorithm
func SearchWithKMP(s, p string) int {
	var i, j, k, r int
	n := len(s)
	m := len(p)
	d := [Mmax]int{}
	k = -1
	d[0] = -1
	for j < m-1 {
		for k >= 0 && p[j] != p[k] {
			k = d[k]
		}
		j++
		k++
		if p[j] == p[k] {
			d[j] = d[k]
		} else {
			d[j] = k
		}
	}
	j = 0
	k = 0
	for j < m && i < n {
		for j >= 0 && s[i] != p[j] {
			j = d[j]
		}
		i++
		j++
	}
	if j == m {
		r = i - m
	} else {
		r = -1
	}
	return r
}
