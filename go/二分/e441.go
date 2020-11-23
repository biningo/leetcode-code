package 二分

/**
*@Author icepan
*@Date 2020/11/23 下午7:08
*@Describe
**/
func arrangeCoins(n int) int {
	c := 0
	i := 1
	for n > 0 {
		if (n - i) < 0 {
			return c
		}
		c++
		n -= i
		i++
	}
	return c
}

//
func arrangeCoins(n int) int {
	left, right := 1, n
	for left <= right {
		mid := (left + right) / 2
		cost := ((mid + 1) * mid) / 2
		if cost == n {
			return mid
		} else if cost > n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
