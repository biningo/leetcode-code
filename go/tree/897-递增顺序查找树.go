package tree

/**
*@Author icepan
*@Date 2020/10/16 下午8:47
*@Describe
**/

var cur *TreeNode = &TreeNode{}

func increasingBST(root *TreeNode) *TreeNode {
	ans := cur
	dfs(root)
	return ans.Right
}

func dfs(node *TreeNode) {
	if node == nil {
		return
	}
	dfs(node.Left)
	node.Left = nil
	cur.Right = node
	cur = node
	dfs(node.Right)
}
