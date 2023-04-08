package main

import (
	"fmt"
)

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	ans := []byte{}
	k := 1
	for n != 0 {
		if n%2 != 0 {
			ans = append(ans, '1')
			n -= k
		} else {
			ans = append(ans, '0')
		}
		k *= -1
		n /= 2
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}
func main() {
	fmt.Println(baseNeg2(2))
	fmt.Println(baseNeg2(3))
	fmt.Println(baseNeg2(4))
}
