package main

import "fmt"

func judge(n int) bool {
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	if sum%2 == 0 {
		return true
	} else {
		return false
	}
}

func countEven(num int) int {
	count := 0
	for i := 2; i <= num; i++ {
		if judge(i) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countEven(30))
	fmt.Println(countEven(4))
}
