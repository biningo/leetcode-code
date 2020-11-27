package 二分

/**
*@Author icepan
*@Date 2020/11/24 下午7:29
*@Describe
**/

func countNegatives(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	r, c := 0, cols-1
	ans := 0
	for c >= 0 && r < rows {
		if grid[r][c] < 0 {
			ans += (rows - r)
			c--
		} else {
			r++
		}
	}
	return ans
}

//
