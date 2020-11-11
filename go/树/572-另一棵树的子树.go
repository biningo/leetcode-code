package 树

/**
*@Author icepan
*@Date 2020/10/16 下午3:52
*@Describe
**/

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return check(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func check(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if (t1 == nil && t2 != nil) || (t1 != nil && t2 == nil) {
		return false
	}
	return check(t1.Left, t2.Left) && check(t1.Right, t2.Right)
}
