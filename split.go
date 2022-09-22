package piscine

func Split(s, sep string) []string {
	slice := 0
	array := []string{}
	for index := range s {
		if index < len(s)-len(sep) {
			if s[index:index+len(sep)] == sep {
				array = append(array, s[slice:index])
				slice = index + len(sep)
			}
		}
		if index == len(s)-1 {
			array = append(array, s[slice:])
		}
	}
	return array
}
