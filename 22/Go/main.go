package main

import "fmt"

func generateParenthesis(n int) (res []string) {

	var backtracking func(n, lp, rp int, path string)
	backtracking = func(n, lp, rp int, path string) {
		if len(path) == 2*n {
			res = append(res, path)
			return
		}
		if lp < n {
			backtracking(n, lp+1, rp, path+"(")
		}
		if rp < lp {
			backtracking(n, lp, rp+1, path+")")
		}
	}
	backtracking(n, 0, 0, "")

	return
}
func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
}
