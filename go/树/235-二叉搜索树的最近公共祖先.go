package 树

/**
*@Author icepan
*@Date 2020/10/8 下午9:51
*@Describe
**/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	node := root
	for {
		if p.Val > node.Val && q.Val > node.Val {
			node = node.Right
		} else if p.Val < node.Val && q.Val < node.Val {
			node = node.Left
		} else { //找到交点了
			return node
		}
	}
}
