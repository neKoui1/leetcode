package main

import "fmt"

func decrypt(code []int, k int) []int {
	if k == 0 {
		return make([]int, len(code))
	}
	temp := append(code, code...)
	if k > 0 {
		for i := 0; i < len(code); i++ {
			sum := 0
			for j := 1; j <= k; j++ {
				sum += temp[i+j]
			}
			temp[i] = sum
		}
		return temp[0:len(code)]
	}
	if k < 0 {
		k = -k
		res := make([]int, 0)
		for i := len(code); i < len(temp); i++ {
			cur := 0
			for j := 1; j <= k; j++ {
				cur += temp[i-j]
			}
			res = append(res, cur)
		}
		return res
	}
	return nil
}
func main() {
	fmt.Println(decrypt([]int{5, 7, 1, 4}, 3))
	fmt.Println(decrypt([]int{1, 2, 3, 4}, 0))
	fmt.Println(decrypt([]int{2, 4, 9, 3}, -2))
}
