package tree

import "strconv"

/**
*@Author icepan
*@Date 2020/10/8 下午9:51
*@Describe
**/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if p.Val > node.Val && q.Val > node.Val {
			queue = append(queue, node.Right)
		} else if p.Val < node.Val && q.Val < node.Val {
			queue = append(queue, node.Left)
		} else { //找到交点了
			return node
		}
	}
	strconv.Itoa()
	return root
}
