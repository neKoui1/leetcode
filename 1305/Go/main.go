package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) (res []int) {
	var Traverse func(root *TreeNode)
	Traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		Traverse(root.Left)
		Traverse(root.Right)
	}

	Traverse(root1)
	Traverse(root2)
	sort.Ints(res)
	return res
}
