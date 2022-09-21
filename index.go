package piscine

func Index(s string, toFind string) int {
	for index := range s {
		if index+len(toFind) > len(s) {
			continue
		}
		if s[index:index+len(toFind)] == toFind {
			return index
		}
	}
	return -1
}
