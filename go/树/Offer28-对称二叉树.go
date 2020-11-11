package 树

/**
*@Author icepan
*@Date 2020/10/17 上午9:47
*@Describe
**/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root.Left, root.Right)
}

func dfs(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}
	if left.Val == right.Val {
		return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}
	return false
}
