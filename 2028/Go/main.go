package main

import "fmt"

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	sum := 0
	for _, v := range rolls {
		sum += v
	}
	target := mean*(m+n) - sum
	temp := target / n
	if temp > 6 || temp <= 0 || temp+target%n <= 0 {
		return nil
	}
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, temp)
	}

	for i := 0; i < target%n; i++ {
		if res[i]+1 > 6 {
			res = nil
			break
		} else {
			res[i]++
		}
	}
	return res
}

func main() {
	rolls := []int{1, 5, 6}
	fmt.Println(missingRolls(rolls, 3, 4))
	rolls = []int{3, 5, 3}
	fmt.Println(missingRolls(rolls, 5, 3))
	rolls = []int{6, 3, 4, 3, 5, 3}
	fmt.Println(missingRolls(rolls, 1, 6))
	rolls = []int{4, 5, 6, 2, 3, 6, 5, 4, 6, 4, 5, 1, 6, 3, 1, 4, 5, 5, 3, 2, 3, 5, 3, 2, 1, 5, 4, 3, 5, 1, 5}
	fmt.Println(missingRolls(rolls, 4, 40))
	rolls = []int{3, 2, 4, 1, 1, 5, 3, 1, 5, 6, 6, 1, 2, 1, 2, 3, 5, 3, 3, 6, 1, 1, 4, 6, 6, 4, 2, 5, 3, 4, 4, 6, 3, 5, 6, 4, 5, 6, 3, 1, 4, 6, 4, 2, 1, 6, 3, 6, 1}
	fmt.Println(missingRolls(rolls, 5, 58))
}
