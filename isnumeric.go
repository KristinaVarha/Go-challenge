package piscine

func IsNumeric(str string) bool {
	h := []rune(str)
	for i := 0; i <= len(h)-1; i++ {
		if (h[i] >= '0') && (h[i] <= '9') {
			return true
		}
	}
	return false
}
