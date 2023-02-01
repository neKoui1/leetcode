package main

import (
	"fmt"
	"strconv"
)

func minimizeXor(num1 int, num2 int) int {
	str1 := strconv.FormatInt(int64(num1), 2)
	str2 := strconv.FormatInt(int64(num2), 2)
	cnt1 := 0
	cnt2 := 0
	for _, v := range str1 {
		if v == '1' {
			cnt1++
		}
	}
	for _, v := range str2 {
		if v == '1' {
			cnt2++
		}
	}
	if cnt1 == cnt2 {
		return num1
	} else if cnt1 < cnt2 {
		min := 0xffffffff
		res := 0
		for num2 != 0 {
			if num2^num1 < min && num2^num1 != 0 {
				min = num2 ^ num1
				res = num2
			}
			num2 >>= 1
		}
		return res
	} else {
		cnt := 0
		str := make([]byte, len(str1))
		for i, v := range str1 {
			str[i] = byte(v)
		}
		for i, v := range str {
			if v == '1' {
				cnt++
				str[i] = '0'
			}
			if cnt == cnt2 {
				break
			}
		}
		res, _ := strconv.ParseInt(string(str), 2, 0)
		return int(res)
	}
}
func main() {
	fmt.Println(minimizeXor(3, 5))
	fmt.Println(minimizeXor(1, 12))
	fmt.Println(minimizeXor(7, 5))
}
