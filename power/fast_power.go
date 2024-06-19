package algdatstr

// The cost is O(log N) where N = n
func FastPowerOf(x float64, n int) float64 {
	var result float64 = 1.0
	i := n
	for {
		if i <= 0 {
			break
		}
		if i%2 != 0 {
			result = result * x
		}
		x = x * x
		i = i / 2
	}
	return result
}
