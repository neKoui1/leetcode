package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	max := 0
	sum := 0
	for _, val := range accounts {
		for _, v := range val {
			sum += v
		}
		if sum > max {
			max = sum
		}
		sum = 0
	}

	return max
}

func main() {
	accounts := [][]int{
		{1, 5},
		{7, 3},
		{3, 5},
	}
	fmt.Println(maximumWealth(accounts))
}
