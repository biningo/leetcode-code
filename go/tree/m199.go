package tree

/**
*@Author icepan
*@Date 2020/11/8 上午10:17
*@Describe
**/

func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		t := []int{}
		for i := 0; i < size; i++ {
			n := queue[0]
			queue = queue[1:]
			t = append(t, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		if len(t) > 0 {
			ans = append(ans, t[len(t)-1])
		}
	}
	return ans
}
