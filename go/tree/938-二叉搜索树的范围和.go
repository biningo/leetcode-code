package tree

/**
*@Author icepan
*@Date 2020/10/16 下午5:14
*@Describe
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	ans := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}

		if node.Val >= L && node.Val <= R {
			ans += node.Val
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else if node.Val < L {
			queue = append(queue, node.Right)
		} else if node.Val > R {
			queue = append(queue, node.Left)
		}
	}
	return ans
}
