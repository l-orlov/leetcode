package climb_stairs

func climbStairs(n int) int {
	res := 1
	tmp := 1
	for i := 1; i < n; i++ {
		res, tmp = res+tmp, res
	}

	return res
}
