package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func binaryGap(n int) int {
	first := 0
	second := 0
	cnt := 0
	pos := 0
	gap := 0
	for n != 0 {
		pos++
		if n&1 == 1 && cnt == 0 {
			first = pos
			cnt++
		} else if n&1 == 1 && cnt == 1 {
			second = pos
			cnt++
		}

		if cnt%2 == 0 && cnt != 0 {
			cnt = 1
			gap = max(gap, second-first)
			first = second
		}

		n >>= 1
	}
	return gap
}
func main() {
	fmt.Println(binaryGap(22))
	fmt.Println(binaryGap(8))
	fmt.Println(binaryGap(5))
}
