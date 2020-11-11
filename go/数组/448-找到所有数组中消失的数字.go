package array

import "math"

/**
*@Author icepan
*@Date 2020/9/24 下午12:19
*@Describe
**/
//将存在的数的下标对应的值变为负数 最后查看没有负数的值
func findDisappearedNumbers(nums []int) []int {
	ans := []int{}
	for _, v := range nums {
		val := int(math.Abs(float64(v)))
		if nums[val-1] > 0 {
			nums[val-1] *= -1
		}
	}
	for i := range nums {
		if nums[i] > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
