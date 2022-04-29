package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	s := fmt.Sprintf("%d", x)
	res := ""
	if s[0] == '-' {
		res += "-"
		s = s[1:]
	}
	for len(s) != 0 {
		res += string(s[len(s)-1])
		s = s[:len(s)-1]
	}
	ans, err := strconv.Atoi(res)
	if err != nil {
		fmt.Println("err =", err)
	}
	if ans > 0x7fffffff || ans < -0x80000000 {
		return 0
	}
	return ans
}
func main() {
	fmt.Println(reverse(-123))
	fmt.Println(reverse(123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(-120))
	fmt.Println(reverse(0))
	fmt.Println(reverse(1534236469))

}
