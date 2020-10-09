package tree

import "strconv"

/**
*@Author icepan
*@Date 2020/10/8 下午10:39
*@Describe
**/
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	ans := []string{}
	dfs2(root, &ans, "")
	return ans
}

func dfs2(node *TreeNode, arr *[]string, s string) {
	if node.Left == nil && node.Right == nil {
		*arr = append(*arr, s+strconv.Itoa(node.Val))
		return
	}
	if node.Left != nil {
		dfs(node.Left, arr, s+strconv.Itoa(node.Val)+"->")
	}
	if node.Right != nil {
		dfs(node.Right, arr, s+strconv.Itoa(node.Val)+"->")
	}
}
