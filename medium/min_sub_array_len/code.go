package min_sub_array_len

import "math"

func minSubArrayLen(target int, nums []int) int {
	minLen := math.MaxInt32
	l, total := 0, 0
	for r := range nums {
		total += nums[r]
		for total >= target {
			tmpLen := r - l + 1
			if tmpLen < minLen {
				minLen = tmpLen
			}
			total -= nums[l]
			l++
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}

	return minLen
}
