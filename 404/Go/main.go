package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		res += root.Left.Val
	}
	return res + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
