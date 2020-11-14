package 链表

/**
*@Author icepan
*@Date 2020/11/14 上午9:34
*@Describe
**/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 {
		return head
	}
	pre := &ListNode{}
	pre.Next = head //考虑一个节点的情况
	fast := pre
	slow := pre
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return pre.Next
}
