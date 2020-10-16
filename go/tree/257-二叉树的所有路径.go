package tree

import "strconv"

/**
*@Author icepan
*@Date 2020/10/8 下午10:39
*@Describe
**/
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	ans := []string{}
	dfs2(root, &ans, "")
	return ans
}

func dfs2(node *TreeNode, arr *[]string, s string) {
	if node.Left == nil && node.Right == nil {
		*arr = append(*arr, s+strconv.Itoa(node.Val))
		return
	}
	if node.Left != nil {
		dfs2(node.Left, arr, s+strconv.Itoa(node.Val)+"->")
	}
	if node.Right != nil {
		dfs2(node.Right, arr, s+strconv.Itoa(node.Val)+"->")
	}
}

type Tuple2 struct {
	Node *TreeNode
	S    string
}

//bfs
func binaryTreePathsBFS(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	queue := []Tuple2{Tuple2{root, strconv.Itoa(root.Val)}}
	ans := []string{}
	for len(queue) > 0 {
		tuple := queue[0]
		queue = queue[1:]
		node, s := tuple.Node, tuple.S
		ss := ""
		if node == root {
			ss = s
		} else {
			ss = s + "->" + strconv.Itoa(node.Val)
		}
		if node.Left == nil && node.Right == nil {
			ans = append(ans, ss)
		}
		if node.Left != nil {
			queue = append(queue, Tuple2{node.Left, ss})
		}
		if node.Right != nil {
			queue = append(queue, Tuple2{node.Right, ss})
		}
	}
	return ans
}
