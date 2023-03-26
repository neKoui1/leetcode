package main

import "fmt"

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	res := 0
	cnt := 0
	for i := 0; i < numOnes && cnt < k; i++ {
		res++
		cnt++
	}
	for i := 0; i < numZeros && cnt < k; i++ {
		cnt++
	}
	for i := 0; i < numNegOnes && cnt < k; i++ {
		res--
		cnt++
	}
	return res

}
func main() {
	fmt.Println(kItemsWithMaximumSum(3, 2, 0, 2))
	fmt.Println(kItemsWithMaximumSum(3, 2, 0, 4))
}
