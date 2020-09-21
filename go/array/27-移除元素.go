package array

/**
*@Author icepan
*@Date 2020/9/20 下午9:46
*@Describe
**/

//1、暴力
func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

//2、双指针
func removeElement2(nums []int, val int) int {
	i, n := 0, len(nums)
	for i < n {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}
	return n
}
