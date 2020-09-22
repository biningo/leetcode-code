package array

import "math"

/**
*@Author icepan
*@Date 2020/9/22 上午8:23
*@Describe
**/

func maxSubArray(nums []int) int {
	ans := -math.MaxInt64
	s := 0
	for _, v := range nums {
		if s > 0 {
			s += v
		} else {
			s = v
		}
		ans = int(math.Max(float64(ans), float64(s)))
	}
	return ans
}
