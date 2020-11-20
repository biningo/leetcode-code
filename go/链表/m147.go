package 链表

/**
*@Author icepan
*@Date 2020/11/17 下午10:42
*@Describe
**/
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := &ListNode{0, nil}
	pre := p
	p.Next = head
	cur := head
	for cur != nil && cur.Next != nil {
		n := cur.Next
		if n.Val < cur.Val {
			cur_next := n.Next
			cur.Next = nil
			find := p.Next
			pre = p
			for find != nil && n.Val > find.Val {
				pre = find
				find = find.Next
			}
			pre.Next = n
			n.Next = find
			cur.Next = cur_next
		} else {
			cur = cur.Next
		}
	}
	return p.Next
}
