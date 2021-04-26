package squaresum

func SquareSum(numbers []int) int {
	i := 0
	sum := 0
	for i < len(numbers) {
		sum += numbers[i] * numbers[i]
		i++
	}
	return sum
}
