package 链表

/**
*@Author icepan
*@Date 2020/11/19 下午8:22
*@Describe
**/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := []int{}
	s2 := []int{}
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	d := 0
	var head *ListNode
	for len(s1) > 0 || len(s2) > 0 || d > 0 {
		n1 := 0
		n2 := 0
		if len(s1) > 0 {
			n1 = s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			n2 = s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		s := (n1 + n2 + d)
		cur := &ListNode{s % 10, nil}
		cur.Next = head
		head = cur
		d = s / 10
	}
	return head
}
