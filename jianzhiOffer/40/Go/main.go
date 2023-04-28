package main

import "sort"

func getLeastNumbers(arr []int, k int) []int {
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr[:k]
}
func main() {

}
