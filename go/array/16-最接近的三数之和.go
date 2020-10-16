package array

import "sort"

/**
*@Author icepan
*@Date 2020/10/4 下午6:36
*@Describe
**/

func threeSumClosest(nums []int, target int) int {
	size := len(nums)
	m := 2 * 1000
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	L := 0
	R := 0
	s := 0
	for i := 0; i < size; i++ {
		L = i + 1
		R = size - 1
		for L < R {
			s = nums[i] + nums[L] + nums[R]
			if abs(target-s) < m {
				ans = s
				m = abs(target - s)
			}
			if s > target {
				R--
			} else if s < target {
				L++
			} else {
				return ans
			}
		}
	}
	return ans
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
