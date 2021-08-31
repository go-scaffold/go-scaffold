package collections

func Sequence(count int) []int {
	if count < 0 {
		count = 0
	}
	result := make([]int, count)
	for i := 0; i < count; i++ {
		result[i] = i
	}
	return result
}
