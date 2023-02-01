package main

//func repeatedNTimes(nums []int) int {
//	hash := make(map[int]int)
//	for _, v := range nums {
//		hash[v]++
//	}
//	n := len(nums) / 2
//	for i, v := range hash {
//		if v == n {
//			return i
//		}
//	}
//	return 0
//}
func repeatedNTimes(nums []int) int {
	hash := make(map[int]int)
	for _, v := range nums {
		hash[v]++
		if hash[v] > 1 {
			return v
		}
	}
	return 0
}
func main() {

}
