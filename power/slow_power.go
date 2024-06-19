package power

// The cost is O(N) where N = n
func SlowPowerOf(x float64, n int) float64 {
	var result float64 = 1.0
	for {
		if n <= 0 {
			break
		}
		result = result * x
		n--
	}
	return result
}
