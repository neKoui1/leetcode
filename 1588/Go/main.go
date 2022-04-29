package main

import (
	"fmt"
)

func sumOddLengthSubarrays(arr []int) int {
	l := len(arr)
	sum := 0
	if l%2 != 0 {
		for l != -1 {
			for i := 0; i < len(arr)-l+1; i++ {
				for j := i; j < l+i; j++ {
					sum += arr[j]
				}
			}
			l -= 2
		}
	} else {
		l--
		for l != -1 {
			for i := 0; i < len(arr)-l+1; i++ {
				for j := i; j < l+i; j++ {
					sum += arr[j]
				}
			}
			l -= 2
		}
	}

	return sum
}

func main() {
	arr := []int{1, 4, 2, 5, 3}
	fmt.Println(sumOddLengthSubarrays(arr))
	arr2 := []int{1, 2}
	fmt.Println(sumOddLengthSubarrays(arr2))
}
