package main

import "fmt"

func climbStairs(n int) (res int) {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	pre, prePre := 2, 1
	for i := 3; i <= n; i++ {
		res = pre + prePre
		pre, prePre = res, pre
	}
	return
}
func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}
