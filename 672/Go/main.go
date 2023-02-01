package main

import "fmt"

/*
	n >= 3 presses >= 3 时
	固定只有8种情况
*/

func flipLights(n int, presses int) int {
	if n == 1 {
		if presses == 0 {
			return 1
		} else if presses >= 1 {
			return 2
		}
	} else if n == 2 {
		if presses == 0 {
			return 1
		} else if presses == 1 {
			return 3
		} else if presses >= 2 {
			return 4
		}
	} else if n >= 3 && presses < 3 {
		if presses == 0 {
			return 1
		} else if presses == 1 {
			return 4
		} else if presses == 2 {
			return 7
		}
	}
	return 8
}
func main() {
	fmt.Println(flipLights(1, 1))
	fmt.Println(flipLights(2, 1))
	fmt.Println(flipLights(3, 1))
}
