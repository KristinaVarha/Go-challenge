package piscine

func Map(f func(int) bool, a []int) []bool {
	array := make([]bool, len(a))
	for i, element := range a {
		array[i] = f(element)
	}
	return array
}
