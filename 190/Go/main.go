package main

import (
	"fmt"
	"math"
)

func reverseBits(num uint32) (res uint32) {
	s := fmt.Sprintf("%032b", num)
	temp := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	for i := 31; i > -1; i-- {
		if temp[i] == '1' {
			res += uint32(math.Pow(2, float64(31-i)))
		}
	}
	return res
}

func main() {
	fmt.Println(reverseBits(0b00000010100101000001111010011100))
	fmt.Println(reverseBits(0b11111111111111111111111111111101))
}
