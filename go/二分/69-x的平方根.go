package 二分

/**
*@Author icepan
*@Date 2020/10/6 下午8:22
*@Describe
**/

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	low, high, mid := 0, x/2, 0
	for low <= high {
		mid = low + (high-low)>>1
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			high = mid - 1
		} else if mid*mid < x {
			low = mid + 1
		}
	}
	return high
}
