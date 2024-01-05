package jump_game

// всегда прыгать максимально, если не можем максимально прыгнуть на текущем шаге, то прыгать на макс. -1

func canJump(nums []int) bool {
	return canJumpFromCell(nums, 0)
}

func canJumpFromCell(nums []int, i int) bool {
	if i >= len(nums)-1 {
		return true
	}

	for j := nums[i]; j > 0; j-- {
		if canJumpFromCell(nums, i+j) {
			return true
		}
	}

	nums[i] = 0

	return false
}
