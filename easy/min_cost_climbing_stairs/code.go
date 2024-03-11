package min_cost_climbing_stairs

func minCostClimbingStairs(cost []int) int {
	minCost := make([]int, len(cost)+1)
	minCost[0] = 0
	minCost[1] = cost[0]
	for i := 2; i <= len(cost); i++ {
		minCost[i] = min(minCost[i-1], minCost[i-2]) + cost[i-1]
	}

	return min(minCost[len(cost)], minCost[len(cost)-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
