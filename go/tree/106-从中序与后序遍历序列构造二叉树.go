package tree

/**
*@Author icepan
*@Date 2020/11/1 上午9:12
*@Describe
**/

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder)==0{
		return nil
	}
	mid:=index(inorder,postorder[len(postorder)-1])
	root:=&TreeNode{
		postorder[len(postorder)-1],
		buildTree(inorder[:mid],postorder[:mid]),
		buildTree(inorder[mid+1:],postorder[mid:len(postorder)-1]),
	}
	return root
}

func index(arr []int,val int) int {
	for i:=0;i<len(arr);i++{
		if arr[i]==val{
			return i
		}
	}
	return -1
}