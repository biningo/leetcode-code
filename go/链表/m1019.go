package 链表

/**
*@Author icepan
*@Date 2020/11/19 下午8:42
*@Describe
**/
func nextLargerNodes(head *ListNode) []int {
	st := []int{}
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	ans := []int{}
	for i := 0; i < len(arr); i++ {
		for len(st) > 0 && arr[i] > arr[st[len(st)-1]] {
			pos := st[len(st)-1]
			st = st[0 : len(st)-1]
			ans[pos] = arr[i]
		}
		ans = append(ans, arr[i])
		st = append(st, i)
	}
	for len(st) > 0 {
		pos := st[len(st)-1]
		st = st[0 : len(st)-1]
		ans[pos] = 0
	}
	return ans
}
