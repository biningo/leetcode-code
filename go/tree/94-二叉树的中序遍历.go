package tree

/**
*@Author icepan
*@Date 2020/10/23 上午9:28
*@Describe
**/

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	stack := []*TreeNode{}
	node := root
	for node != nil {
		stack = append(stack, node)
		node = node.Left
	}

	for len(stack) > 0 {
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		node = node.Right
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
	}
	return ans
}
