package main

import "fmt"

func countOperations(num1 int, num2 int) (res int) {
	for num1 > 0 && num2 > 0 {
		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
		res++
	}
	return
}
func main() {
	fmt.Println(countOperations(10, 10))
	fmt.Println(countOperations(2, 3))
	fmt.Println(countOperations(0, 0))
}
