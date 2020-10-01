package array

/**
*@Author icepan
*@Date 2020/10/1 上午8:58
*@Describe
**/

func countNegatives(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}

	row, col := 0, cols-1
	ans := 0
	for col >= 0 && row < rows {
		if grid[row][col] < 0 {
			ans += (rows - row)
			col--
		} else {
			row++
		}
	}
	return ans
}
