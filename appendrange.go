package piscine

func AppendRange(min, max int) []int {
	size := max - min
	var answer []int
	if size <= 0 {
		return nil
	}
	for i := min; i < max; i++ {
		answer = append(answer, i)
	}
	return answer
}
