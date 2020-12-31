package uncertainty

func Within(f, target, e float64) bool {
	if f < (target + e) {
		if f > (target - e) {
			return true
		}
	}
	return false
}
