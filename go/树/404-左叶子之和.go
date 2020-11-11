package 树

/**
*@Author icepan
*@Date 2020/10/9 下午6:58
*@Describe
**/

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	if root.Left != nil {
		if isLeaf(root.Left) == true {
			ans = ans + root.Left.Val
		} else {
			ans = ans + sumOfLeftLeaves(root.Left)
		}
	}
	if root.Right != nil {
		ans += sumOfLeftLeaves(root.Right)
	}
	return ans
}

func isLeaf(node *TreeNode) bool {
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}

//------------------------------------BFS------------------------
func sumOfLeftLeavesBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			if isLeaf(node.Left) {
				ans += node.Left.Val
			} else {
				queue = append(queue, node.Left)
			}
		}
		if node.Right != nil {
			if isLeaf(node.Right) == false {
				queue = append(queue, node.Right)
			}
		}
	}
	return ans
}
