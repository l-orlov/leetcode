package summary_ranges

import "fmt"

func summaryRanges(nums []int) []string {
	var res []string

	start := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] != nums[i]+1 {
			if i == start {
				res = append(res, fmt.Sprintf("%d", nums[i]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[i]))
			}
			start = i + 1
		}
	}

	if start < len(nums) {
		if start == len(nums)-1 {
			res = append(res, fmt.Sprintf("%d", nums[start]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[len(nums)-1]))
		}
	}

	return res
}
