package array

import "math"

/**
*@Author icepan
*@Date 2020/9/22 上午8:54
*@Describe
**/

func maxProfit(prices []int) int {
	maxprofit := 0
	minprice := math.MaxInt64
	for _, v := range prices {
		maxprofit = max(maxprofit, v-minprice)
		minprice = min(minprice, v)
	}
	return maxprofit
}
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
