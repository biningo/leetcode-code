package tree

/**
*@Author icepan
*@Date 2020/10/23 上午9:28
*@Describe
**/

//中序
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	stack := []*TreeNode{}
	node := root
	for node != nil {
		stack = append(stack, node)
		node = node.Left
	}

	for len(stack) > 0 {
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		node = node.Right
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
	}
	return ans
}


//前序
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	stack := []*TreeNode{}
	node := root
	for node != nil {
		stack = append(stack, node)
		ans = append(ans, node.Val)
		node = node.Left
	}

	for len(stack) > 0 {
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
		for node != nil {
			ans = append(ans, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
	}
	return ans
}


//后序
func postorderTraversal(root *TreeNode) []int {
	if root==nil{
		return nil
	}
	stack:=[]*TreeNode{}
	ans:=[]int{}
	pre:=&TreeNode{}
	node:=root
	for node!=nil{
		stack = append(stack,node)
		node = node.Left
	}
	for len(stack)>0{
		node:=stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right==nil || node.Right==pre{
			ans = append(ans,node.Val)
			pre = node
			node = nil
		}else{
			stack = append(stack,node)
			node = node.Right
		}
		for node!=nil{
			stack = append(stack,node)
			node = node.Left
		}
	}
	return ans
}