package max_product_diff

import (
	"sort"
)

func maxProductDifference(nums []int) int {
	if len(nums) < 4 {
		return -1
	}

	firstNums := []int{nums[0], nums[1], nums[2], nums[3]}
	sort.Slice(firstNums, func(i, j int) bool {
		return firstNums[i] > firstNums[j]
	})

	maxVal, maxVal2 := firstNums[0], firstNums[1]
	minVal2, minVal := firstNums[2], firstNums[3]

	for i := 4; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal2 = maxVal
			maxVal = nums[i]
		} else if nums[i] < minVal {
			minVal2 = minVal
			minVal = nums[i]
		} else if nums[i] > maxVal2 {
			maxVal2 = nums[i]
		} else if nums[i] < minVal2 {
			minVal2 = nums[i]
		}
	}

	return (maxVal * maxVal2) - (minVal * minVal2)
}
