package 二分

/**
*@Author icepan
*@Date 2020/11/23 下午2:52
*@Describe
**/

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return []int{-1, -1}
	}
	size := len(numbers)
	for i := 0; i < size; i++ {
		left, right := i+1, size-1
		for left <= right {
			mid := (left + right) / 2
			s := numbers[i] + numbers[mid]
			if s == target {
				return []int{i + 1, mid + 1}
			} else if s > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return []int{-1, -1}
}
