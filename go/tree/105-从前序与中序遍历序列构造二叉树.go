package tree

/**
*@Author icepan
*@Date 2020/10/25 下午9:44
*@Describe
**/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder)==0{
		return nil
	}
	mid:=index(inorder,preorder[0])
	root:=&TreeNode{
		preorder[0],
		buildTree(preorder[1:mid+1],inorder[:mid]),
		buildTree(preorder[mid+1:],inorder[mid+1:]),
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