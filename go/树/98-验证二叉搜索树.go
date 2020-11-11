package 树

/**
*@Author icepan
*@Date 2020/10/24 下午3:10
*@Describe
**/
//中序遍历解决
func isValidBST(root *TreeNode) bool {
	pre:=-(1<<63 - 1)
	return dfs(&pre,root)
}

func dfs(pre *int,node *TreeNode) bool{
	if node==nil{
		return true
	}
	if dfs(pre,node.Left)==false{
		return false
	}
	if node.Val<=*pre{
		return false
	}
	*pre = node.Val
	return dfs(pre,node.Right)
}