package piscine

func ForEach(f func(int), a []int) {
	for _, results := range a {
		f(results)
	}
}
