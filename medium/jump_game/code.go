package jump_game

// всегда прыгать максимально, если не можем максимально прыгнуть на текущем шаге, то прыгать на макс. -1

func canJump(nums []int) bool {
	canJumpIdx := 0
	for i, num := range nums {
		if i > canJumpIdx {
			return false
		}
		canJumpIdx = max(canJumpIdx, i+num)
	}

	return true
}

func jump2(nums []int) int {
	rightMaxPos, maxPos := 0, 0
	answer := 0
	for i := 0; i < len(nums)-1 && maxPos < len(nums); i++ {
		rightMaxPos = max(rightMaxPos, i+nums[i])
		if i == maxPos {
			answer++
			maxPos = rightMaxPos
		}
	}

	return answer
}
