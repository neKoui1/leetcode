package main

import "fmt"

func smallestRepunitDivByK(k int) int {
	p := 0
	for i := 1; i <= k; i++ {
		p = (p*10 + 1) % k
		if p == 0 {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Println(smallestRepunitDivByK(1))
	fmt.Println(smallestRepunitDivByK(2))
	fmt.Println(smallestRepunitDivByK(2))
}
