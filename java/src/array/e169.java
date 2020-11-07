package array;

/***
 *@Author icepan
 *@Date 2020/11/7 上午9:35
 *@Description
 *
 ***/


public class e169 {
    public int majorityElement(int[] nums) {
        int vote=0;
        int ans = nums[0];
        for(int i=0;i<nums.length;i++){
            if(vote==0){
                ans = nums[i];
                vote=1;
            }else if(nums[i]==ans){
                vote++;
            }else{
                vote--;
            }
        }
        return ans;
    }
}
