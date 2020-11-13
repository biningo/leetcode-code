package 链表

/**
*@Author icepan
*@Date 2020/11/12 下午12:41
*@Describe
**/

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode = nil
	current := head
	for current != nil {
		next_node := current.Next
		current.Next = pre
		pre = current
		current = next_node
	}
	return pre
}
