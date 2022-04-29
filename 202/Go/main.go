package main

import "fmt"

func isHappy(n int) bool {
	sum := 0
	temp := 0
	for n != 0 {
		temp = n % 10
		n /= 10
		sum += temp * temp
		if n < 10 {
			sum += n * n
			if sum == 1 {
				return true
			} else if sum == 4 {
				break
			}
			n = sum
			sum = 0
		}
	}
	return false
}

func main() {
	fmt.Println(isHappy(3))
}
