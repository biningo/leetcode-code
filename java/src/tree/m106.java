package tree;

import java.util.HashMap;
import java.util.Map;

/***
 *@Author icepan
 *@Date 2020/11/1 上午8:46
 *@Description
 *
 ***/


public class m106 {
    private Map<Integer,Integer> indexMap;

    public TreeNode helper(int[] postorder,int[] inorder,int p_left,int p_right,int i_left,int i_right){
        if(p_left>p_right){
            return null;
        }
        int mid = indexMap.get(postorder[p_right]);
        int subLength = mid-i_left;
        TreeNode root = new TreeNode(postorder[p_right]);
        root.left = helper(postorder,inorder,p_left,p_left+subLength-1,i_left,mid-1);
        root.right = helper(postorder,inorder,p_left+subLength,p_right-1,mid+1,i_right);
        return root;
    }


    public TreeNode buildTree(int[] inorder, int[] postorder) {
        indexMap = new HashMap<>();
        for (int i = 0; i < inorder.length; i++) {
            indexMap.put(inorder[i],i);
        }
        return helper(postorder,inorder,0,postorder.length-1,0,inorder.length-1);
    }
}
