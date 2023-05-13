package main

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func numMovesStones(a int, b int, c int) []int {
	x := Min(a, Min(b, c))
	z := Max(a, Max(b, c))
	y := a + b + c - x - z
	min, max := 0, 0
	if z-x > 2 {
		min = 2
		if y-x < 3 || z-y < 3 {
			min = 1
		}
		max = z - x - 2
	}
	return []int{min, max}
}
func main() {

}
