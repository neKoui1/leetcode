package main

import (
	"fmt"
	"strconv"
)

func evenOddBit(n int) (res []int) {
	str := strconv.FormatInt(int64(n), 2)
	cnt := 0
	evenCnt := 0
	oddCnt := 0
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '1' {
			if cnt%2 == 0 {
				evenCnt++
			} else {
				oddCnt++
			}
		}
		cnt++
	}
	res = append(res, evenCnt)
	res = append(res, oddCnt)
	return
}
func main() {
	fmt.Println(evenOddBit(17))
	fmt.Println(evenOddBit(2))
}
