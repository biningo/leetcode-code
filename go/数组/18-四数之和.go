package array

import "sort"

/**
*@Author icepan
*@Date 2020/10/3 上午10:34
*@Describe
**/

func fourSum(nums []int, target int) [][]int {
	size := len(nums)
	if size < 4 {
		return [][]int{}
	}
	sort.Ints(nums)
	s := 0
	L := 0
	R := 0
	ans := [][]int{}
	for i := 0; i < size-3; i++ {
		//去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//最小的都比target大 break
		if (nums[i] + nums[i+1] + nums[i+2] + nums[i+3]) > target {
			break
		}
		//最大的都比target小 continue
		if (nums[i] + nums[size-1] + nums[size-2] + nums[size-3]) < target {
			continue
		}

		//三数之和
		for j := i + 1; j < size-2; j++ {
			if j > (i+1) && nums[j] == nums[j-1] {
				continue
			}

			if (nums[j] + nums[j+1] + nums[j+2]) > (target - nums[i]) {
				break
			}
			if (nums[j] + nums[size-1] + nums[size-2]) < (target - nums[i]) {
				continue
			}

			L = j + 1
			R = size - 1

			for L < R {
				s = nums[j] + nums[L] + nums[R]
				if s == (target - nums[i]) {
					ans = append(ans, []int{nums[i], nums[j], nums[L], nums[R]})
					for L < R && nums[L] == nums[L+1] {
						L++
					}
					for L < R && nums[R] == nums[R-1] {
						R--
					}
					R--
					L++
				} else if s > (target - nums[i]) {
					R--
				} else if s < (target - nums[i]) {
					L++
				}
			}
		}
	}
	return ans
}
