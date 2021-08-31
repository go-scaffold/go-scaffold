package arithmetics

func Sum(values ...int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}
