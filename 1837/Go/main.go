package main

import (
	"fmt"
	"strconv"
)

func sumBase(n int, k int) int {
	s := strconv.FormatInt(int64(n), k)
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	res := 0
	for num != 0 {
		res += num % 10
		num /= 10
	}
	return res
}
func main() {
	fmt.Println(sumBase(34, 6))
}
