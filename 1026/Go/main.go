package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// node.val >= 0
func maxAncestorDiff(root *TreeNode) (res int) {
	stack := make([]int, 0)

	var dfs func(node *TreeNode, nums []int)
	dfs = func(node *TreeNode, nums []int) {
		if node == nil {
			M := math.MinInt
			m := math.MaxInt
			for _, v := range nums {
				if v > M {
					M = v
				}
				if v < m {
					m = v
				}
				res = Max(res, M-m)
			}
			return
		}

		dfs(node.Left, append(append([]int{}, nums...), node.Val))
		dfs(node.Right, append(append([]int{}, nums...), node.Val))
	}
	dfs(root, stack)
	return res
}
func main() {

}
