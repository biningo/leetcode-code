package 二分

/**
*@Author icepan
*@Date 2020/11/24 下午7:38
*@Describe
**/

//二分查找
func minArray(numbers []int) int {
	low, hight := 0, len(numbers)-1
	for low < hight {
		mid := low + (hight-low)/2
		if numbers[mid] > numbers[hight] {
			low = mid + 1
		} else if numbers[mid] < numbers[hight] {
			hight = mid
		} else {
			hight--
		}
	}
	return numbers[hight]
}

//遍历
func minArray(numbers []int) int {
	size := len(numbers)
	if size == 1 {
		return numbers[0]
	}
	i := 0
	for (i + 1) < size {
		if numbers[i+1] < numbers[i] {
			return numbers[i+1]
		}
		i++
	}
	return numbers[0]
}
