package 链表

/**
*@Author icepan
*@Date 2020/11/13 上午9:32
*@Describe
**/

//1、栈
func getKthFromEnd(head *ListNode, k int) *ListNode {
	stack := []*ListNode{}
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}
	ans := &ListNode{}
	for i := 0; i < k; i++ {
		ans = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return ans
}

//2、双指针
func getKthFromEnd2(head *ListNode, k int) *ListNode {
	p1 := head
	p2 := head
	for i := 0; i < k; i++ {
		if p2.Next == nil {
			return p1
		}
		p2 = p2.Next
	}
	for p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
