package 链表

/**
*@Author icepan
*@Date 2020/11/15 下午3:35
*@Describe
**/

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	var pre *ListNode = nil
	c := n - m + 1 //要反转的节点个数
	for m > 1 {
		pre = cur
		cur = cur.Next
		m--
	}
	//此时cur指向待反转链表的第一个节点 pre指向前一个
	h, t := pre, cur //t指向反转之后的最后一个节点
	for c > 0 {
		next_node := cur.Next
		cur.Next = pre
		pre = cur
		cur = next_node
		c--
	}
	//此时cur指向最后一个节点的后一个节点
	t.Next = cur
	//预防从第一个节点开始旋转的例子
	if h != nil {
		h.Next = pre
	} else {
		head = pre
	}
	return head
}
