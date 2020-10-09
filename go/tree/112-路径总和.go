package tree

/**
*@Author icepan
*@Date 2020/10/8 下午12:55
*@Describe
**/

type Tuple struct {
	Node *TreeNode
	S    int
}

func hasPathSumDFS(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	stack := []Tuple{Tuple{root, sum - root.Val}}
	for len(stack) > 0 {
		size := len(stack)
		tuple := stack[size-1]
		stack = stack[:size-1]
		node, s := tuple.Node, tuple.S
		if node.Left == nil && node.Right == nil && s == 0 {
			return true
		}
		if node.Left != nil {
			stack = append(stack, Tuple{node.Left, s - node.Left.Val})
		}
		if node.Right != nil {
			stack = append(stack, Tuple{node.Right, s - node.Right.Val})
		}
	}
	return false
}

func hasPathSumBFS(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	queue := []Tuple{Tuple{root, sum - root.Val}}
	for len(queue) > 0 {
		tuple := queue[0]
		queue = queue[1:]
		node, s := tuple.Node, tuple.S
		if node.Left == nil && node.Right == nil && s == 0 {
			return true
		}
		if node.Left != nil {
			queue = append(queue, Tuple{node.Left, s - node.Left.Val})
		}
		if node.Right != nil {
			queue = append(queue, Tuple{node.Right, s - node.Right.Val})
		}
	}
	return false
}
