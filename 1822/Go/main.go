package main

import "fmt"

func signFunc(product int) int {
	if product > 0 {
		return 1
	} else if product < 0 {
		return -1
	} else {
		return 0
	}
}

func arraySign(nums []int) int {
	product := 1
	for _, v := range nums {
		product *= v
		if product > 0 {
			product = 1
		} else if product < 0 {
			product = -1
		} else {
			return 0
		}
	}
	return signFunc(product)
}

func main() {
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1}))
	fmt.Println(arraySign([]int{1, 5, 0, 2, -3}))
	fmt.Println(arraySign([]int{-1, 1, -1, 1, -1}))
	fmt.Println(arraySign([]int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}))
}
