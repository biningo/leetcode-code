package array

/**
*@Author icepan
*@Date 2020/9/27 下午5:37
*@Describe
**/

func validMountainArray(A []int) bool {
	length := len(A)
	if length < 3 {
		return false
	}
	left := 0
	right := length - 1
	for left < right {
		if A[left] < A[left+1] {
			left++
		} else if A[right] < A[right-1] {
			right--
		} else {
			break
		}
	}
	return left == right && left != 0 && right != length-1
}
