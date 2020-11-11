package 树

/**
*@Author icepan
*@Date 2020/10/8 上午9:25
*@Describe
**/

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return dfs(0, len(nums)-1, nums)
}
func dfs(low int, high int, nums []int) *TreeNode {
	if low > high {
		return nil
	}
	mid := low + ((high - low) >> 1)
	node := &TreeNode{Val: nums[mid]}
	node.Left = dfs(low, mid-1, nums)
	node.Right = dfs(mid+1, high, nums)
	return node
}
