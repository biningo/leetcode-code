package 链表

/**
*@Author icepan
*@Date 2020/11/16 下午1:41
*@Describe
**/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	var pre *ListNode = nil
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if pre != nil {
		pre.Next = nil
	}
	node := &TreeNode{slow.Val, nil, nil}
	if slow == head { //只有一个节点的情况
		return node
	}
	node.Left = sortedListToBST(head)
	node.Right = sortedListToBST(slow.Next)
	return node
}
