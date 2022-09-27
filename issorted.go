package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	res := 1
	res2 := 1
	res3 := 1
	for index, element := range a {
		if index != len(a)-1 {
			if f(element, a[index+1]) < 0 {
				res++
			}
			if f(element, a[index+1]) > 0 {
				res2++
			}
			if f(element, a[index+1]) == 0 {
				res3++
			}
		}
	}
	return res == len(a) || res2 == len(a) || res3 == len(a)
}
