package main

import (
	"fmt"
	"math"
)

func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}
	x := math.Log(float64(n)) / math.Log(4)
	return x == math.Trunc(x)
}
func main() {
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(5))
	fmt.Println(isPowerOfFour(1))
}
