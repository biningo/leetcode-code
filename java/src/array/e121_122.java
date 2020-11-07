package array;

/***
 *@Author icepan
 *@Date 2020/11/2 下午2:21
 *@Description
 *
 ***/


public class e121_122 {
    public int maxProfit_121(int[] prices) {
        if (prices.length==0){
            return 0;
        }
        int maxprofit = 0;
        int minLeft=prices[0];
        for(int p:prices){
            if(p>minLeft){
                maxprofit = Math.max(maxprofit,p-minLeft);
            }
            minLeft = Math.min(minLeft,p);
        }
        return maxprofit;
    }


    public int maxProfit_122(int[] prices) {
        if(prices.length==0){
            return 0;
        }
        int sumProfit=0;
        int i=0;
        while ((i+1)<prices.length){
            if(prices[i+1]>prices[i]){
                sumProfit +=(prices[i+1]-prices[i]);
            }
            i++;
        }
        return sumProfit;
    }
}
