package 链表

/**
*@Author icepan
*@Date 2020/11/14 上午9:52
*@Describe
**/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	k = k % linkLen(head)
	if k == 0 {
		return head
	}
	slow := head
	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	h := slow.Next
	slow.Next = nil
	fast.Next = head
	return h
}
func linkLen(head *ListNode) int {
	c := 0
	for head != nil {
		c++
		head = head.Next
	}
	return c
}
