package candy

func candy(ratings []int) int {
	length := len(ratings)
	candies := make([]int, length)

	for i := range candies {
		candies[i] = 1
	}

	// left to the right
	for i := 1; i < length; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// left to the right
	for i := length - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	totalCandies := 0
	for _, c := range candies {
		totalCandies += c
	}

	return totalCandies
}
