package 树

/**
*@Author icepan
*@Date 2020/11/7 上午10:43
*@Describe
**/

var ans = 0

func longestUnivaluePath(root *TreeNode) int {
	dfs(root)
	return ans
}
func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftSubLen := dfs(node.Left)
	rightSubLen := dfs(node.Right)

	leftLen := 0
	rightLen := 0
	if node.Left != nil && node.Left.Val == node.Val {
		leftLen = leftSubLen + 1
	}
	if node.Right != nil && node.Right.Val == node.Val {
		rightLen = rightSubLen + 1
	}
	ans = max(ans, leftLen+rightLen)
	return max(leftLen, rightLen)
}

func max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
