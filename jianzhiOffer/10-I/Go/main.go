package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % (1e9 + 7)
	}
	return r
}
func main() {
	fmt.Println(fib(5))
	fmt.Println(fib(95))
}
