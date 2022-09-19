package piscine

func IterativePower(nb int, power int) int {
	var result int = 1
	if power > 0 {
		result = nb * (IterativePower(nb, power-1))
	} else {
		return 0
	}
	return result
}
