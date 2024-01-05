package spiral_matrix

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	result := make([]int, 0, len(matrix)*len(matrix[0]))
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for {
		// right
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}

		// down
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if left > right {
			break
		}

		// left
		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}

		// выход
		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return result
}
