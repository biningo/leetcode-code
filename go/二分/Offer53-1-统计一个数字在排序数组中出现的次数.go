package 二分

/**
*@Author icepan
*@Date 2020/10/3 下午5:56
*@Describe
**/

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	mid := 0
	size := len(nums)
	//寻找右边界
	for low <= high {
		mid = low + ((high - low) >> 1)
		if nums[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if high >= 0 && nums[high] != target {
		return 0
	}

	right := low

	//寻找左边界
	low = 0
	for low <= high {
		mid = low + ((high - low) >> 1)
		if nums[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if low < size && nums[low] != target {
		return 0
	}
	left := high
	return right - left - 1
}
