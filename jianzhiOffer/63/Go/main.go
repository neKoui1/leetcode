package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	profix := 0
	buyPrice := math.MaxInt
	for _, price := range prices {
		if buyPrice > price {
			buyPrice = price
		}
		if price-buyPrice > profix {
			profix = price - buyPrice
		}
	}
	return profix
}
func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
