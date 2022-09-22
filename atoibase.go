package piscine

func AtoiBase(s string, base string) int {
	res := 0
	lens := 0
	arr := map[rune]bool{}
	for _, c := range base {
		if arr[c] || c == '-' || c == '+' {
			return res
		}
		arr[c] = true
		lens++
	}
	if lens > 1 {
		for _, v := range s {
			count := 0
			if arr[v] {
				for _, j := range base {
					if j == v {
						break
					}
					count++
				}
				res = res*lens + count
			}
		}
	}
	return res
}
