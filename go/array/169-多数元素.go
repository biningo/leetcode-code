package array

/**
*@Author icepan
*@Date 2020/9/22 下午10:14
*@Describe
**/

func majorityElement(nums []int) int {
	val:=0
	vote:=0
	for _,v:=range nums{
		if vote==0{
			val=v
		}
		if val==v{
			vote++
		}else{
			vote--
		}
	}
	return val
}