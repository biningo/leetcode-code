package array;

/***
 *@Author icepan
 *@Date 2020/10/27 下午6:42
 *@Description
 *
 ***/


public class e26_27 {

    //26
    public int removeDuplicates(int[] nums) {
        if(nums.length==0){
            return 0;
        }
        int i=0;
        for(int j=i+1;j<nums.length;j++){
            if(nums[i]!=nums[j]){
                i++;
                nums[i] = nums[j];
            }
        }
        return i+1;
    }

    //27
    public int removeElement(int[] nums, int val) {
        if(nums.length==0){
            return 0;
        }
        int i=0;
        for(int j=0;j<nums.length;j++){
            if(nums[j]!=val){
                nums[i] = nums[j];
                i++;
            }
        }
        return i;
    }
}
