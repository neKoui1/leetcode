package main

import "fmt"

func cross(p, q, r []int) int {
	return (q[0]-p[0])*(r[1]-q[1]) - (q[1]-p[1])*(r[0]-q[0])
}

func outerTrees(trees [][]int) [][]int {
	res := make([][]int, 0)
	n := len(trees)
	if n < 4 {
		return trees
	}

	leftMost := 0
	for i, tree := range trees {
		if tree[0] < trees[leftMost][0] {
			leftMost = i
		}
	}

	vis := make([]bool, n)
	p := leftMost
	for {
		q := (p + 1) % n
		for r, tree := range trees {
			if cross(trees[p], trees[q], tree) < 0 {
				q = r
			}
		}

		for i, b := range vis {
			if !b && i != p && i != q && cross(trees[p], trees[q], trees[i]) == 0 {
				res = append(res, trees[i])
				vis[i] = true
			}
		}
		if !vis[q] {
			res = append(res, trees[q])
			vis[q] = true
		}
		p = q
		if p == leftMost {
			break
		}
	}

	return res
}
func main() {
	trees := [][]int{
		{1, 1},
		{2, 2},
		{2, 0},
		{2, 4},
		{3, 3},
		{4, 2},
	}
	fmt.Println(outerTrees(trees))
}
