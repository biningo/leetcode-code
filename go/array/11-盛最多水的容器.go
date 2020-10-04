package array

/**
*@Author icepan
*@Date 2020/10/3 上午9:26
*@Describe
**/

func maxArea(height []int) int {
	size := len(height)
	left, right := 0, size-1
	ans := (right - left) * min(height[left], height[right])

	for left < right {
		if height[left] < height[right] {
			left++
			ans = max(ans, (right-left)*min(height[left], height[right]))
		} else {
			right--
			ans = max(ans, (right-left)*min(height[left], height[right]))
		}
	}
	return ans
}
