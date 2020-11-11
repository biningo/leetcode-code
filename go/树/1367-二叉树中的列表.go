package 树

/**
*@Author icepan
*@Date 2020/11/11 下午12:50
*@Describe
**/

func isSubPath(head *ListNode, root *TreeNode) bool {
	if head==nil{
		return true
	}
	if root==nil{
		return false
	}
	return isSub(head,root) || isSubPath(head,root.Left) || isSubPath(head,root.Right)
}

func isSub(list_node *ListNode,tree_node *TreeNode) bool{
	if list_node==nil{
		return true
	}
	if tree_node==nil{
		return false
	}
	if list_node.Val!=tree_node.Val{
		return false
	}
	return isSub(list_node.Next,tree_node.Left) || isSub(list_node.Next,tree_node.Right)
}