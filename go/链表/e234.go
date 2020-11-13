package 链表

/**
*@Author icepan
*@Date 2020/11/13 上午8:59
*@Describe
**/

func isPalindrome(head *ListNode) bool {
	//1、快慢指针寻找中间节点
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//2、反转后面的链表
	var pre *ListNode = nil
	current := slow
	for current != nil {
		next_node := current.Next
		current.Next = pre
		pre = current
		current = next_node
	}
	//3、比较两条链表是否相等
	h1 := pre
	h2 := head
	for h1 != nil && h2 != nil {
		if h1.Val != h2.Val {
			return false
		}
		h1 = h1.Next
		h2 = h2.Next
	}
	return true
}
