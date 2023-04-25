package main

import "fmt"

func getNum(n int) int {
	res := 0
	for n != 0 {
		res += n % 10
		n /= 10
	}
	return res
}

func movingCount(m int, n int, k int) int {
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	var moving func(i, j int) int
	moving = func(i, j int) int {
		if i >= m || j >= n || i < 0 || j < 0 || getNum(i)+getNum(j) > k {
			return 0
		}
		if vis[i][j] {
			return 0
		}
		vis[i][j] = true
		return moving(i+1, j) + moving(i, j+1) + 1
	}

	return moving(0, 0)
}

func main() {
	fmt.Println(movingCount(2, 3, 1))
	fmt.Println(movingCount(3, 1, 0))
}
