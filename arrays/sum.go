package arrays

// Sum takes an array of length five and returns the sum of those numbers
func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// SumAll calculates sum for slice lists
func SumAll(numbersList ...[]int) []int {
	var sums []int
	for _, nums := range numbersList {
		sums = append(sums, Sum(nums))
	}
	return sums
}
