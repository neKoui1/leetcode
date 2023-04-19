package main

import "fmt"

func numWays(n int) int {
	const mod int = 1e9 + 7
	pre, prePre := 1, 1
	for i := 2; i <= n; i++ {
		prePre, pre = pre%mod, (pre+prePre)%mod
	}
	return pre
}
func main() {
	fmt.Println(numWays(2))
	fmt.Println(numWays(7))
	fmt.Println(numWays(0))
}
