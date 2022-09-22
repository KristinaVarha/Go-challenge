package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	var cache string
	for _, char := range s {
		if IsPrintable(string(char)) && char != ' ' {
			cache += string(char)
		} else if len(cache) > 0 {
			result = append(result, cache)
			cache = ""
		}
	}
	if len(cache) > 0 {
		result = append(result, cache)
	}
	return result
}
