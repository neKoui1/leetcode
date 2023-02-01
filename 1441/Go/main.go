package main

import "fmt"

func buildArray(target []int, n int) (res []string) {
	list := make([]int, 0)
	flag := 0
	for i := 1; i <= n; i++ {
		list = append(list, i)
	}
	for _, v := range list {
		if v == target[flag] {
			res = append(res, "Push")
			flag++
			if flag == len(target) {
				break
			}
		} else {
			res = append(res, "Push")
			res = append(res, "Pop")
			continue
		}
	}
	return
}
func main() {
	fmt.Println(buildArray([]int{1, 3}, 3))
	fmt.Println(buildArray([]int{1, 2, 3}, 3))
	fmt.Println(buildArray([]int{1, 2}, 4))
}
