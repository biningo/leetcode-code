package 链表

/**
*@Author icepan
*@Date 2020/11/18 上午9:34
*@Describe
**/

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMid(head)
	left, right := sortList(head), sortList(mid)
	node := &ListNode{0, nil}
	res := node
	for left != nil && right != nil {
		if left.Val < right.Val {
			node.Next = left
			left = left.Next
		} else {
			node.Next = right
			right = right.Next
		}
		node = node.Next
	}
	if left != nil {
		node.Next = left
	} else {
		node.Next = right
	}
	return res.Next
}

func findMid(node *ListNode) *ListNode {
	slow, fast := node, node.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	return mid
}
