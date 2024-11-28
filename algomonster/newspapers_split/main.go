package newspapers_split

import (
	"math"
)

func feasible(newspapersReadTimes []int, numCoworkers, limit int) bool {
	// time to keep track of the current worker's time spent, numWorkers to keep track of the number of coworkers used
	time, numWorkers := 0, 0
	for _, readTime := range newspapersReadTimes {
		// check if current time exceeds the given time limit
		if time+readTime > limit {
			time = 0
			numWorkers++
		}
		time += readTime
	}
	// edge case to check if we needed an extra worker at the end
	if time != 0 {
		numWorkers++
	}
	// check if the number of workers we need is more than what we have
	return numWorkers <= numCoworkers
}

func newspapersSplit(newspapersReadTimes []int, numCoworkers int) int {
	// Initializing 'low' to the maximum integer value and 'high' to 0
	low, high := math.MinInt32, 0
	// Finding the minimum value in the array and accumulating the total sum
	for _, readTime := range newspapersReadTimes {
		low = int(math.Max(float64(low), float64(readTime)))
		high += readTime
	}
	// Using binary search to find the optimal time limit
	ans := -1
	for low <= high {
		mid := (low + high) / 2
		// helper function to check if a time works
		if feasible(newspapersReadTimes, numCoworkers, mid) {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}
