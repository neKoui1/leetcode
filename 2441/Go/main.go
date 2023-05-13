package main

type Set map[int]struct{}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// hash or two pointers
func findMaxK(nums []int) int {
	set := make(Set)
	for _, v := range nums {
		set[v] = struct{}{}
	}
	ans := -1
	for _, v := range nums {
		if v < 0 {
			continue
		}
		if _, ok := set[-v]; ok {
			ans = Max(ans, v)
		}
	}
	return ans
}
func main() {

}
