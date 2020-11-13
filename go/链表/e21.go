package 链表

/**
*@Author icepan
*@Date 2020/11/11 下午5:06
*@Describe
**/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{0, nil}
	h := l3
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l3.Next = l1
			l1 = l1.Next
		} else {
			l3.Next = l2
			l2 = l2.Next
		}
		l3 = l3.Next
	}
	if l1 != nil {
		l3.Next = l1
	} else {
		l3.Next = l2
	}
	return h.Next
}
