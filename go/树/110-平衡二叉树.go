package 树

/**
*@Author icepan
*@Date 2020/10/8 上午9:59
*@Describe
**/

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(dfs(root.Left)-dfs(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(dfs(node.Left), dfs(node.Right)) + 1
}
func max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
