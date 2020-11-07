package array;

/***
 *@Author icepan
 *@Date 2020/10/26 下午10:54
 *@Description
 *
 ***/


public class e7 {
    public int reverse(int x) {
        int ans=0;
        while (x!=0){
            int pop = x%10;
            //溢出直接返回0
            if(ans>(Integer.MAX_VALUE/10) || (ans==(Integer.MAX_VALUE/10)&&pop>7)){
                return 0;
            }
            if(ans<(Integer.MIN_VALUE/10) || (ans==(Integer.MIN_VALUE/10)&&pop<-8)){
                return 0;
            }
            ans = ans*10+pop;
            x /=10;
        }
        return ans;
    }
}
