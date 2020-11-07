package tree;


import java.util.HashMap;
import java.util.Map;

/***
 *@Author icepan
 *@Date 2020/10/31 上午9:08
 *@Description
 *
 ***/


public class m105 {

    private Map<Integer,Integer> indexMap;
    public TreeNode helper(int[] preorder,int[] inorder,int p_left,int p_right,int i_left,int i_right){
        if(p_left>p_right){
            return null;
        }
        TreeNode root = new TreeNode(preorder[p_left]);
        int mid = indexMap.get(preorder[p_left]);
        int subLength =  mid-i_left;
        root.left = helper(preorder,inorder,p_left+1,p_left+subLength,i_left,mid-1);
        root.right = helper(preorder,inorder,p_left+subLength+1,p_right,mid+1,i_right);
        return root;
    }

    public TreeNode buildTree(int[] preorder, int[] inorder) {
        indexMap = new HashMap<>();
        for (int i = 0; i < inorder.length; i++) {
            indexMap.put(inorder[i],i);
        }
        return helper(preorder,inorder,0,preorder.length-1,0,inorder.length-1);
    }





}
