package three_sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var result [][]int
	var tmpTarget int
	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		for v := i + 1; v < len(nums)-2; v++ {
			if v != i+1 && nums[v] == nums[v-1] {
				continue
			}

			tmpTarget = target - nums[i] - nums[v]
			for j, k := v+1, len(nums)-1; j < k; {
				if nums[j]+nums[k] == tmpTarget {
					result = append(result, []int{nums[i], nums[v], nums[j], nums[k]})
					j++
					for j < k && nums[j] == nums[j-1] {
						j++
					}
				} else if nums[j]+nums[k] < tmpTarget {
					j++
				} else {
					k--
				}
			}
		}
	}

	return result
}
