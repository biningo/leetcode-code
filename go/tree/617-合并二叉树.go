package tree

/**
*@Author icepan
*@Date 2020/10/10 上午9:11
*@Describe
**/

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

type TupleTree struct {
	N1 *TreeNode
	N2 *TreeNode
}

func mergeTreesBFS(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	queue := []TupleTree{TupleTree{t1, t2}}
	for len(queue) > 0 {
		val := queue[0]
		queue = queue[1:]
		n1, n2 := val.N1, val.N2
		if n2 == nil { //n1肯定不为nil
			continue
		}
		n1.Val += n2.Val
		if n1.Left == nil {
			n1.Left = n2.Left
		} else {
			queue = append(queue, TupleTree{n1.Left, n2.Left})
		}
		if n1.Right == nil {
			n1.Right = n2.Right
		} else {
			queue = append(queue, TupleTree{n1.Right, n2.Right})
		}
	}
	return t1
}
