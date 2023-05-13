package main

func numPairsDivisibleBy60(time []int) int {
	ans := 0
	record := make([]int, 61)
	for _, v := range time {
		x := v % 60
		y := (60 - x) % 60
		ans += record[y]
		record[x]++
	}

	return ans
}
func main() {

}
