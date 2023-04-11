package main

import (
	"fmt"
)

func dailyTemperatures(temperatures []int) (res []int) {
	monoStack := make([]int, 0)
	n := len(temperatures)
	res = make([]int, n)
	for i := 0; i < n; i++ {

		for len(monoStack) != 0 && temperatures[monoStack[len(monoStack)-1]] < temperatures[i] {
			tempIdx := monoStack[len(monoStack)-1]
			monoStack = monoStack[:len(monoStack)-1]
			res[tempIdx] = i - tempIdx
		}
		monoStack = append(monoStack, i)
	}
	return
}
func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))
}
