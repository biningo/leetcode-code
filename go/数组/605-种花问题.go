package array

/**
*@Author icepan
*@Date 2020/9/25 上午8:49
*@Describe
**/

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	for i := 0; i < length; i++ {
		if n == 0 {
			return true
		}
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == (length-1) || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
	}
	if n == 0 {
		return true
	}
	return false
}
