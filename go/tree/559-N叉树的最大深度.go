package tree

/**
*@Author icepan
*@Date 2020/10/9 下午7:19
*@Describe
**/
type Node struct {
	Val      int
	Children []*Node
}

func maxDepthN(root *Node) int {
	if root == nil {
		return 0
	}
	ans := 0
	queue := []*Node{root}
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			for i := 0; i < len(node.Children); i++ {
				queue = append(queue, node.Children[i])
			}
			size--
		}
		ans++
	}
	return ans
}

//DFS
func maxDepthDFS(root *Node) int {
	if root == nil {
		return 0
	}
	ans := 0
	for i := 0; i < len(root.Children); i++ {
		ans = max(ans, maxDepthDFS(root.Children[i]))
	}
	return ans + 1
}
