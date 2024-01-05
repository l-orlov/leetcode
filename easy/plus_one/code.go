package plus_one

func plusOne(digits []int) []int {
	i := len(digits) - 1
	digits[i]++
	for digits[i] == 10 && i > 0 {
		digits[i] = 0
		i--
		digits[i]++
	}

	if i == 0 && digits[i] == 10 {
		digits[i] = 0
		digits = append([]int{1}, digits...)
	}

	return digits
}
