package piscine

func Join(strs []string, sep string) string {
	res := ""
	n := len(strs)
	for i, a := range strs {
		res += a
		if i < n-1 {
			res += sep
		}
	}
	return res
}
