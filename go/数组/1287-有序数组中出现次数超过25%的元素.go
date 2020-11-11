package array

/**
*@Author icepan
*@Date 2020/10/1 上午8:37
*@Describe
**/

func findSpecialInteger(arr []int) int {
	length := len(arr)
	c := 0
	cur := arr[0]

	for i := 0; i < length; i++ {
		if arr[i] == cur {
			c++
			if c*4 > length {
				return cur
			}
		} else {
			c = 1
			cur = arr[i]
		}
	}
	return cur
}
