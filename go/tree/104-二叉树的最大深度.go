package tree

/**
*@Author icepan
*@Date 2020/10/7 下午4:27
*@Describe
**/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	size := 0
	ans := 0
	for len(queue) > 0 {
		size = len(queue)
		for size > 0 {
			n := queue[0]
			queue = queue[1:]
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
			size--
		}
		//遍历为一层 +1
		ans++
	}
	return ans
}
