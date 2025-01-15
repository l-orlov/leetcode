package happy_number

func isHappy(n int) bool {
	var seenNums = map[int]struct{}{}

	for n != 1 {
		n = getSumOfSquares(n)
		if _, ok := seenNums[n]; ok {
			return false
		}
		seenNums[n] = struct{}{}
	}

	return true
}

func getSumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}

	return sum
}
