package 树

/**
*@Author icepan
*@Date 2020/10/7 下午3:52
*@Describe
**/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//bfs
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	queue1, queue2 := []*TreeNode{p}, []*TreeNode{q}

	for len(queue1) > 0 && len(queue2) > 0 {
		n1, n2 := queue1[0], queue2[0]
		if n1.Val != n2.Val {
			return false
		}

		queue1, queue2 = queue1[1:], queue2[1:]
		left1, left2 := n1.Left, n2.Left
		right1, right2 := n1.Right, n2.Right

		//如果两个树的节点有一个为nil 另外一个为真 则返回false
		//两个同时为true 或则 nil 则继续判断
		if (left1 == nil && left2 != nil) || (left1 != nil && left2 == nil) {
			return false
		}
		if (right1 == nil && right2 != nil) || (right1 != nil && right2 == nil) {
			return false
		}

		if left1 != nil {
			queue1 = append(queue1, left1)
		}
		if right1 != nil {
			queue1 = append(queue1, right1)
		}
		if left2 != nil {
			queue2 = append(queue2, left2)
		}
		if right2 != nil {
			queue2 = append(queue2, right2)
		}
	}
	return len(queue1) == 0 && len(queue2) == 0
}
