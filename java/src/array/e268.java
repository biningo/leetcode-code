package array;

/***
 *@Author icepan
 *@Date 2020/11/7 上午9:45
 *@Description
 *
 ***/


public class e268 {
    public int missingNumber(int[] nums) {
        int ans = nums.length;
        for(int i=0;i<nums.length;i++){
            ans^=(i^nums[i]);
        }
        return ans;
    }
}
