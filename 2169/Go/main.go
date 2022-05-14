package main

import "fmt"

/*
	也算是一种递归罢
*/
func countOperations(num1 int, num2 int) (cnt int) {
	var recur func(num1, num2 int)
	recur = func(num1, num2 int) {
		if num1 == 0 || num2 == 0 {
			return
		}
		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
		cnt++
		recur(num1, num2)
	}
	recur(num1, num2)
	return
}
func main() {
	fmt.Println(countOperations(2, 3))
	fmt.Println(countOperations(10, 10))
	fmt.Println(countOperations(0, 0))
}
