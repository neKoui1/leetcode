package main

import "fmt"

func Abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) (res int) {
	for _, v1 := range arr1 {
		flag := true
		for _, v2 := range arr2 {
			if Abs(v1, v2) <= d {
				flag = false
				break
			}
		}
		if flag {
			res++
		}
	}
	return
}

func main() {
	arr1 := []int{4, 5, 8}
	arr2 := []int{10, 9, 1, 8}
	fmt.Println(findTheDistanceValue(arr1, arr2, 2))

	a := []int{1, 4, 2, 3}
	b := []int{-4, -3, 6, 10, 20, 30}
	fmt.Println(findTheDistanceValue(a, b, 3))

	a = []int{2, 1, 100, 3}
	b = []int{-5, -2, 10, -3, 7}
	fmt.Println(findTheDistanceValue(a, b, 6))
}
