package array

/**
*@Author icepan
*@Date 2020/9/29 上午8:21
*@Describe
**/

func prefixesDivBy5(A []int) []bool {
	t := 0
	ans := make([]bool, len(A))
	for i, v := range A {
		t = t*2 + v
		if t%5 == 0 {
			ans[i] = true
		} else {
			ans[i] = false
		}
		t %= 10 //防止超过最大整数
	}
	return ans
}
