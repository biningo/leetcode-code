package tree

/**
*@Author icepan
*@Date 2020/10/16 下午4:54
*@Describe
**/

func minDiffInBST(root *TreeNode) int {

	arr := []int{}
	dfs(&arr, root)
	ans := 101
	i := 0
	for (i + 1) < len(arr) {
		ans = min(ans, abs(arr[i]-arr[i+1]))
		i++
	}
	return ans
}
func dfs(arr *[]int, node *TreeNode) {
	if node == nil {
		return
	}
	dfs(arr, node.Left)
	*arr = append(*arr, node.Val)
	dfs(arr, node.Right)
}

func min(v1 int, v2 int) int {
	if v1 > v2 {
		return v2
	}
	return v1
}
func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
