import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class test {
    public static void main(String[] args){
        List<Integer> list = new ArrayList<>();
        list.add(1);
        list.add(2);
        int[] num = new int[list.size()];
        for(int i = 0;i < list.size();i++){
            num[i] = list.get(i);
        }
        for(int i = 0; i < list.size(); i++){
            System.out.print(num[i]);
        }
    }
}
