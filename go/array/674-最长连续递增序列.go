package array

/**
*@Author icepan
*@Date 2020/9/26 上午8:40
*@Describe
**/

func findLengthOfLCIS(nums []int) int {
	ans := 0
	t := 0
	length := len(nums)
	i := 0
	for i < length {
		t = 0
		for i < length-1 && nums[i+1] > nums[i] {
			t++
			i++
		}
		i++
		ans = max(ans, t+1)
	}
	return ans
}
