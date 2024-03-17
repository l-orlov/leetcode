package number_of_1_bits

func hammingWeight(num uint32) int {
	count := 0

	for i := num; i != 0; {
		i = i & (i - 1)
		count++
	}

	return count
}
