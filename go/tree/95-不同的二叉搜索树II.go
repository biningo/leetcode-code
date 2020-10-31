package tree

/**
*@Author icepan
*@Date 2020/10/24 下午3:57
*@Describe
**/

func generateTrees(n int) []*TreeNode {
	if n==0{
		return []*TreeNode{}
	}
	return dfs(1,n)
}

func dfs(start int,end int) []*TreeNode{
	if start>end{
		return []*TreeNode{nil} //需要返回一个nil
	}
	ans:=[]*TreeNode{}
	for i:=start;i<=end;i++{
		leftTrees:=dfs(start,i-1)
		rightTrees:=dfs(i+1,end)
		for l:=0;l<len(leftTrees);l++{
			for r:=0;r<len(rightTrees);r++{
				node:=&TreeNode{i,leftTrees[l],rightTrees[r]}
				ans =append(ans,node)
			}
		}
	}
	return ans
}