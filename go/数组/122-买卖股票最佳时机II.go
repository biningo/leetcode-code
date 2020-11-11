package array

/**
*@Author icepan
*@Date 2020/9/22 上午9:01
*@Describe
**/

func maxProfit2(prices []int) int {
	profit := 0
	n := len(prices)
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			profit += (prices[i] - prices[i-1])
		}
	}
	return profit
}
