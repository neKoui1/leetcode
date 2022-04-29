package main

import "fmt"

func addToArrayForm(num []int, k int) []int {
	OF := 0
	for i := len(num) - 1; i != -1; i-- {
		temp := k % 10
		k /= 10
		num[i] = num[i] + temp + OF
		OF = num[i] / 10
		num[i] %= 10
	}
	for k != 0 {
		k += OF
		temp := k % 10
		k /= 10
		OF = temp / 10
		num = append([]int{temp}, num...)
	}
	if OF > 0 {
		num = append([]int{OF}, num...)
	}
	return num
}

func main() {
	fmt.Println(addToArrayForm([]int{1, 2, 0, 0}, 34))
	fmt.Println(addToArrayForm([]int{2, 7, 4}, 181))
	fmt.Println(addToArrayForm([]int{2, 1, 5}, 806))
	fmt.Println(addToArrayForm([]int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3}, 516))
	fmt.Println(addToArrayForm([]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1))
	fmt.Println(addToArrayForm([]int{0}, 23))
	fmt.Println(addToArrayForm([]int{6}, 809))
}
