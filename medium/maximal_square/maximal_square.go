package maximal_square

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	dp := make([]int, n+1)
	maxSize, prev := 0, 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			temp := dp[j]

			if matrix[i-1][j-1] == '1' {
				dp[j] = min(dp[j], dp[j-1], prev) + 1
				maxSize = max(maxSize, dp[j])
			} else {
				dp[j] = 0
			}

			prev = temp
		}
	}

	return maxSize * maxSize
}
