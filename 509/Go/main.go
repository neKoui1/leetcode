package main

import "fmt"

func fib(n int) (res int) {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	pre := 1
	prePre := 0
	for i := 2; i <= n; i++ {
		res = pre + prePre
		pre, prePre = res, pre
	}
	return
}
func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
}
