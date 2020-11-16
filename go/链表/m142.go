package 链表

/**
*@Author icepan
*@Date 2020/11/16 下午1:55
*@Describe
**/

//快慢指针
func detectCycle2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { //最终肯定会在环中相遇
			break
		}
	}
	//无环的情况
	if fast == nil || fast.Next == nil {
		return nil
	}

	//继续走
	p := head
	for p != slow {
		p = p.Next
		slow = slow.Next
	}
	return p
}

//哈希
func detectCycle(head *ListNode) *ListNode {
	p := head
	m := make(map[*ListNode]int)
	for p != nil {
		if m[p] == 1 {
			return p
		}
		m[p] = 1
		p = p.Next
	}
	return nil
}
