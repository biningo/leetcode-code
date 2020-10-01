package array

/**
*@Author icepan
*@Date 2020/9/30 上午8:57
*@Describe
**/

func relativeSortArray(arr1 []int, arr2 []int) []int {
	bucket := make([]int, 1001)
	for _, v := range arr1 {
		bucket[v]++
	}

	cur := 0
	for _, v := range arr2 {
		for bucket[v] > 0 {
			arr1[cur] = v
			bucket[v]--
			cur++
		}
	}

	for v := range bucket {
		for bucket[v] > 0 {
			arr1[cur] = v
			cur++
			bucket[v]--
		}
	}
	return arr1
}
