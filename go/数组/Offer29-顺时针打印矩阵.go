package array

/**
*@Author icepan
*@Date 2020/10/4 下午7:29
*@Describe
**/

func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0 {
		return []int{}
	}
	cols := len(matrix[0])
	if cols == 0 {
		return []int{}
	}

	visited := make([][]bool, rows)
	for r := 0; r < rows; r++ {
		visited[r] = make([]bool, cols)
	}
	count := 0
	ans := make([]int, rows*cols)

	col := -1
	row := 0
	for count < rows*cols {
		for (col+1) < cols && !visited[row][col+1] {
			col += 1
			ans[count] = matrix[row][col]
			visited[row][col] = true
			count++
		}

		for (row+1) < rows && !visited[row+1][col] {
			row += 1
			ans[count] = matrix[row][col]
			visited[row][col] = true
			count++
		}

		for (col-1) >= 0 && !visited[row][col-1] {
			col -= 1
			ans[count] = matrix[row][col]
			visited[row][col] = true
			count++
		}

		for (row-1) >= 0 && !visited[row-1][col] {
			row -= 1
			ans[count] = matrix[row][col]
			visited[row][col] = true
			count++
		}
	}
	return ans
}
