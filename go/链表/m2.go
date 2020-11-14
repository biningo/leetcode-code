package 链表

/**
*@Author icepan
*@Date 2020/11/13 下午6:25
*@Describe
**/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s := 0
	d := 0
	head := l1
	for l1 != nil && l2 != nil {
		s = (l1.Val + l2.Val + d)
		l1.Val = s % 10
		d = s / 10
		if l1.Next == nil || l2.Next == nil {
			break
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1.Next == nil {
		l1.Next = l2.Next
	}
	for l1.Next != nil && d > 0 {
		l1 = l1.Next
		s = (l1.Val + d)
		l1.Val = s % 10
		d = s / 10
	}
	if d > 0 {
		l1.Next = &ListNode{d, nil}
	}
	return head
}
