package array;

/***
 *@Author icepan
 *@Date 2020/11/7 上午9:58
 *@Description
 *
 ***/


public class e283 {
    public void moveZeroes(int[] nums) {
        for(int i=0;i<nums.length;i++){
            for(int j=0;j<(nums.length-i-1);j++){
                if(nums[j]==0){
                    int t = nums[j+1];
                    nums[j+1]=0;
                    nums[j] = t;
                }
            }
        }
    }
}
