package piscine

func BasicAtoi(s string) int {
	runes := []rune(s)
	num := 0
	length := 0
	for i := range runes {
		length = i
	}
	ten := 1
	for j := 0; j < length; j++ {
		ten *= 10
	}
	for i := range runes {
		temp := 0
		for k := '0'; k < runes[i]; k++ {
			temp++
		}
		num += temp * ten
		ten /= 10
	}
	return num
}
