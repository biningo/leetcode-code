package 二分

/**
*@Author icepan
*@Date 2020/11/23 下午7:19
*@Describe
**/

func peakIndexInMountainArray(arr []int) int {
	i := 0
	for (i + 1) < len(arr) {
		if arr[i+1] < arr[i] {
			return i
		}
		i++
	}
	return i
}

//二分
func peakIndexInMountainArray(arr []int) int {
	lfet, right := 0, len(arr)-1
	for lfet <= right {
		mid := lfet + (right-lfet)/2
		if arr[mid] < arr[mid+1] {
			lfet = mid + 1
		} else {
			right = mid - 1
		}
	}
	return lfet
}
