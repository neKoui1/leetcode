import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Solution {
    public static List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        List<Integer> temp = new ArrayList<>();
        if(nums.length == 1){
            temp.add(nums[0]);
            res.add(temp);
            return res;
        }else if(nums.length == 2){
            temp.add(nums[0]);
            temp.add(nums[1]);
            res.add(temp);
            temp.clear();
            // Java 的 List 有removeAll和clear两种方法清空列表
            // 查菜鸟看下来， clear 的效率比removeAll要高
            temp.add(nums[1]);
            temp.add(nums[0]);
            res.add(temp);
            return res;
        }else{
            List<Integer> tList = new ArrayList<I>(Arrays.asList(nums));
            for(int i = 0; i < nums.length; i++){
                List<Integer> list = tList;
                list.remove(i);
                int[] paraNums = new int[list.size()];
                for(int j = 0; j < list.size(); j++){
                    paraNums[j] = list.get(j);
                }
                List<List<Integer>> tempList = permute(paraNums);
                for(List<Integer> value : tempList){
                    value.add(nums[i]);
                    res.add(value);
                }
            }
            return res;
        }
    }

    public static void main(String[] args){
        int[] nums = {1,2,3};
        System.out.println(permute(nums));
    }
}
