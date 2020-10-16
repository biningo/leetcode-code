package tree

/**
*@Author icepan
*@Date 2020/10/10 上午9:11
*@Describe
**/

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	// L<root.Val<R 合理范围 递归修剪左子树 右子树
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
