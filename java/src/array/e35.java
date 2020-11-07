package array;

/***
 *@Author icepan
 *@Date 2020/10/27 下午7:07
 *@Description
 *
 ***/


public class e35 {
    public int maxSubArray(int[] nums) {
        int m=nums[0];
        int temp=0;
        for(int i=0;i<nums.length;i++){
            if(temp<0){
                temp=0;
            }
            temp += nums[i];
            m = Math.max(m,temp);
        }
        return m;
    }
}
