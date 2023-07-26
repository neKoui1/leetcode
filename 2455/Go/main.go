package main

func averageValue(nums []int) int {
	sum, cnt := 0, 0
	for _, v := range nums {
		if v%3 == 0 && v%2 == 0 {
			sum += v
			cnt++
		}
	}
	if cnt == 0 {
		return 0
	}
	return sum / cnt
}
func main() {

}
