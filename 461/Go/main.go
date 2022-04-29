//两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。
//给你两个整数 x 和 y，计算并返回它们之间的汉明距离。
package main

import "fmt"

func hammingDistance(x int, y int) int {
	count := 0
	s := fmt.Sprintf("%b", x^y)
	for _, v := range s {
		if v == '1' {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(hammingDistance(1, 4))
}
