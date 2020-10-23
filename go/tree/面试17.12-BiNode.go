package tree

/**
*@Author icepan
*@Date 2020/10/23 上午9:26
*@Describe
**/
var cur *TreeNode
var ans *TreeNode

func convertBiNode(root *TreeNode) *TreeNode {
	cur = &TreeNode{0, nil, nil}
	ans = cur
	dfs(root)
	return ans.Right
}
func dfs(node *TreeNode) {
	if node == nil {
		return
	}
	dfs(node.Left)
	cur.Right = node
	node.Left = nil
	cur = node
	dfs(node.Right)
}
