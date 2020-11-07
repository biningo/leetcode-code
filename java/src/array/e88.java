package array;

/***
 *@Author icepan
 *@Date 2020/10/28 下午3:23
 *@Description
 *
 ***/


public class e88 {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int p1,p2,p;
        p1=m-1;
        p2=n-1;
        p=m+n-1;
        while (p1>=0 && p2>=0){

            if(nums1[p1]>nums2[p2]){
                nums1[p] = nums1[p1];
                p1--;
            }else{
                nums1[p] = nums2[p2];
                p2--;
            }
            p--;
        }
        while (p2>=0){
            nums1[p] = nums2[p2];
            p2--;p--;
        }
    }
}
