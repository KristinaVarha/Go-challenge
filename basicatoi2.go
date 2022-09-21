package piscine

func BasicAtoi2(s string) int {
	arr := []rune(s)
	y := 0
	final := 0
	for range s {
		y++
	}
	for x := 0; x <= y-1; x++ {
		if int(arr[x]-'0') >= 0 && int(arr[x]-'0') < 10 {
			final = final*10 + int(arr[x]-'0')
		} else {
			return 0
		}
	}
	return final
}
