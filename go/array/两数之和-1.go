package array

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if index, ok := m[target-v]; ok {
			return []int{index, i}
		}
		m[v] = i
	}
	return []int{-1, -1}
}
