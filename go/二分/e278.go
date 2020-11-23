package 二分

/**
*@Author icepan
*@Date 2020/10/6 下午8:36
*@Describe
**/
func isBadVersion(v int) bool {
	return true
}

//二分
func firstBadVersion(n int) int {
	low, high := 1, n
	mid := 0
	for low <= high {
		mid = (low + high) / 2
		if isBadVersion(mid) == false {
			low = mid + 1
		} else {
			if mid == 1 || isBadVersion(mid-1) == false {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

//2
func firstBadVersion(n int) int {
	left, right := 1, n
	for left <= right {
		mid := (left + right) / 2
		if isBadVersion(mid) == true {
			for isBadVersion(mid) == true {
				mid--
			}
			return mid + 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
