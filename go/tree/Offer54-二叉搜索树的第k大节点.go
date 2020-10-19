package tree

/**
*@Author icepan
*@Date 2020/10/19 下午12:22
*@Describe
**/

func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{root}
	node := root
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Right
		}

		if len(stack) > 0 {
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			k -= 1
			if k == 0 {
				return n.Val
			}
			node = n.Left
		}
	}
	return 0
}
