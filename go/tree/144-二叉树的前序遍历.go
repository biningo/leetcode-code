package tree

/**
*@Author icepan
*@Date 2020/10/31 上午10:28
*@Describe
**/

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	stack := []*TreeNode{}
	node := root
	for node != nil {
		stack = append(stack, node)
		ans = append(ans, node.Val)
		node = node.Left
	}

	for len(stack) > 0 {
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
		for node != nil {
			ans = append(ans, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
	}
	return ans
}
