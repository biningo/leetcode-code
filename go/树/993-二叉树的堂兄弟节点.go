package 树

type Tuple5 struct {
	Pre  *TreeNode
	Node *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	queue := []Tuple5{Tuple5{nil, root}}

	for len(queue) > 0 {
		xExist := Tuple5{nil, nil}
		yExist := Tuple5{nil, nil}
		size := len(queue)
		for i := 0; i < size; i++ {
			t := queue[0]
			queue = queue[1:]
			_, node := t.Pre, t.Node
			if node.Val == x {
				xExist = t
			} else if node.Val == y {
				yExist = t
			}
			if node.Left != nil {
				queue = append(queue, Tuple5{node, node.Left})
			}
			if node.Right != nil {
				queue = append(queue, Tuple5{node, node.Right})
			}
		}
		if xExist.Node != nil && yExist.Node != nil && xExist.Pre != nil && xExist.Pre != yExist.Pre {
			return true
		}
	}
	return false
}

//方法二
func isCousins2(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	dept := make(map[int]int)
	parent := make(map[int]int)
	dept[root.Val] = 0   //val对应深度
	parent[root.Val] = 0 //存放父亲的值

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			dept[node.Left.Val] = dept[node.Val] + 1
			parent[node.Left.Val] = node.Val
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			dept[node.Right.Val] = dept[node.Val] + 1
			parent[node.Right.Val] = node.Val
			queue = append(queue, node.Right)
		}
	}
	return dept[x] == dept[y] && parent[x] != parent[y]
}
