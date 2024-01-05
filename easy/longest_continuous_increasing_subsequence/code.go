package longest_continuous_increasing_subsequence

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxLen, tmpLen := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			tmpLen++
		} else {
			if tmpLen > maxLen {
				maxLen = tmpLen
			}
			tmpLen = 1
		}
	}
	if tmpLen > maxLen {
		maxLen = tmpLen
	}

	return maxLen
}
