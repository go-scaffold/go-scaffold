package arithmetics

func Subtract(values ...int) int {
	result := values[0]
	for i := 1; i < len(values); i++ {
		result -= values[i]
	}
	return result
}
