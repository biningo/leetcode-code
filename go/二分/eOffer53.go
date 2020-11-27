package 二分

/**
*@Author icepan
*@Date 2020/11/25 下午4:16
*@Describe
**/
//二分法
func missingNumber(nums []int) int {
	size := len(nums)
	left, right := 0, size-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= mid { //左边全部没有缺失 直接看右边
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

//位运算
func missingNumber(nums []int) int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		n ^= (i ^ nums[i])
	}
	return n
}
