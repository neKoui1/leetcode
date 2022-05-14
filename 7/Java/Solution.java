class Solution {
    public static int reverse(int x) {
        long res = 0;
        while(x != 0){
            int y = x % 10;
            x  = x / 10;
            res = res * 10 + y;
        }
        return (res > 0x7fffffff || res < 0x80000000)?0:(int)res;
    }

    public static void main(String[] args){
        System.out.println(reverse(100));
        System.out.println(reverse(123));
        System.out.println(reverse(-123));
        System.out.println(reverse(1534236469));
        System.out.println(0x7fffffff);
        System.out.println(0x80000000);
    }
}