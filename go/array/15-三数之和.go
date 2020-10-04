package array

import "sort"

/**
*@Author icepan
*@Date 2020/10/3 上午9:58
*@Describe
**/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	size := len(nums)
	ans := [][]int{}
	L := 0
	R := size - 1
	s := 0

	for i := 0; i < size; i++ {
		if nums[i] > 0 {
			break
		}
		//去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		L = i + 1
		R = size - 1

		for L < R {
			s = nums[i] + nums[L] + nums[R]
			if s == 0 {
				ans = append(ans, []int{nums[i], nums[L], nums[R]})
				//去重
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if s > 0 {
				R--
			} else if s < 0 {
				L++
			}
		}
	}
	return ans
}

//target通用版
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	size := len(nums)
	ans := [][]int{}
	L := 0
	R := size - 1
	s := 0

	target := 0
	for i := 0; i < size-2; i++ {

		//去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//最小的都大于target 直接break
		if (nums[i] + nums[i+1] + nums[i+2]) > target {
			break
		}
		//最大的都小于target continue
		if (nums[i] + nums[size-1] + nums[size-2]) < target {
			continue
		}

		L = i + 1
		R = size - 1

		for L < R {
			s = nums[i] + nums[L] + nums[R]
			if s == target {
				ans = append(ans, []int{nums[i], nums[L], nums[R]})
				//去重
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if s > target {
				R--
			} else if s < target {
				L++
			}
		}
	}
	return ans
}
