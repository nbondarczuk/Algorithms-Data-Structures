package power

import "fmt"

// Is it like [0.]*
func isZeroOrDotStr(str string) bool {
	var i, n int
	n = len(str)
	for {
		if i >= n {
			break
		}
		// char not in {'0', '.'}
		if str[i] != '0' && str[i] != '.' {
			return false
		}
		i++
	}
	return true
}

func IsPowerOf10(x float64) bool {
	if x < 0 {
		return false
	}
	str := fmt.Sprintf("%v", x)
	if x >= 1.0 {
		// Like 1[0.]* when x >= 1.0
		return str[0] == '1' && isZeroOrDotStr(str[1:])
	}
	// Like [0.]*1 when x >= 0 && x <= 1.0
	return str[len(str)-1] == '1' && isZeroOrDotStr(str[:len(str)-1])
}
