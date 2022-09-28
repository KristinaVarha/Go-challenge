package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 1; i < len(a); i++ {
		for j := 1; j < len(a); j++ {
			if f(a[j-1], a[j]) == 1 {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
}
