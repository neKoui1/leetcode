package main

import (
	"fmt"
	"strconv"
)

/*
	我恨遍历
	我恨循环
	这个方法真的好蠢
*/
func complexNumberMultiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	var r1, r2, c1, c2 int
	for i, v := range num1 {
		if v == '+' {
			r1, _ = strconv.Atoi(num1[:i])
			c1, _ = strconv.Atoi(num1[i+1 : m-1])
		}
	}
	for i, v := range num2 {
		if v == '+' {
			r2, _ = strconv.Atoi(num2[:i])
			c2, _ = strconv.Atoi(num2[i+1 : n-1])
		}
	}
	real := r1*r2 - c1*c2
	imag := r1*c2 + r2*c1
	return fmt.Sprintf("%d+%di", real, imag)
}

func main() {
	num1 := "1+1i"
	fmt.Println(num1[2 : len(num1)-1])
	fmt.Println(complexNumberMultiply("1+1i", "1+1i"))
	fmt.Println(complexNumberMultiply("78+-76i", "-86+72i"))
}
