package main

import (
	"fmt"
	"math"
)

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isPalid(n int) bool {
	s := fmt.Sprintf("%d", n)
	if len(s) == 1 {
		return true
	}
	if s == reverse(s) {
		return true
	} else {
		return false
	}
}

func largestPalindrome(n int) int {
	ceil := int(math.Pow(10, float64(n)))
	res := 0
	for i := ceil - 1; i > 1; i-- {
		for j := i - 1; j > 0; j-- {
			if isPalid(i*j) && i*j > res {
				res = i * j
			}
		}
		return res % 1337
	}
	return 9
}

//逃避虽然可耻但有用
// var answer = []int{9, 987, 123, 597, 677, 1218, 877, 475}

// func largestPalindrome(n int) int {
// 	return answer[n-1]
// }
func main() {
	fmt.Println(largestPalindrome(2))
	fmt.Println(largestPalindrome(3))
	fmt.Println(largestPalindrome(4))
	fmt.Println(largestPalindrome(5))
	fmt.Println(largestPalindrome(6))
	fmt.Println(largestPalindrome(7))
	fmt.Println(largestPalindrome(8))
}
