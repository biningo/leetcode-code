package 链表

/**
*@Author icepan
*@Date 2020/11/14 上午9:43
*@Describe
**/

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1 := head
	p2 := head.Next
	p1.Next = swapPairs(p2.Next)
	p2.Next = p1
	return p2
}
