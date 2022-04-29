package main

import "fmt"

func average(salary []int) float64 {
	max := 0
	min := 0xffffffff
	sum := 0
	for _, v := range salary {
		sum += v
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	sum = sum - min - max
	return float64(sum) / float64(len(salary)-2)
}

func main() {
	salary := []int{1000, 2000, 3000}
	fmt.Println(average(salary))
}
