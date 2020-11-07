package array;

import java.util.*;
import java.util.stream.Collectors;
import java.util.stream.Stream;

/***
 *@Author icepan
 *@Date 2020/10/28 下午12:25
 *@Description
 *
 ***/


public class e66 {
    public int[] plusOne(int[] digits) {
        int d=0;
        int s = digits[digits.length-1]+1+d;
        digits[digits.length-1] = s%10;
        d = s/10;
        for (int i = digits.length-2; i >=0 ; i--) {
            if(d==0){
                return digits;
            }
            s = digits[i]+d;
            digits[i] = s%10;
            d = s/10;
        }
        if(d!=0){
            int ans[] = new int[digits.length+1];
            ans[0] = d;
            for (int i = 0; i < digits.length; i++) {
                ans[i+1] = digits[i];
            }
            return ans;
        }
        return digits;
    }

}
