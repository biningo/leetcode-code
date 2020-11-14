package 链表

/**
*@Author icepan
*@Date 2020/11/14 上午10:26
*@Describe
**/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1 := &ListNode{0, nil}
	h := p1
	p1.Next = head

	p2 := head
	p3 := p2.Next

	for p3 != nil {
		if p2.Val != p3.Val {
			p1 = p2
			p2 = p3
			p3 = p2.Next
		} else {
			for p3 != nil && p3.Val == p2.Val {
				p3 = p3.Next
			}
			p1.Next = p3
			p2 = p3
			if p2 != nil {
				p3 = p2.Next
			}
		}
	}
	return h.Next
}
