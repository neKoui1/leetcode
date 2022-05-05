package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (res []int) {
	var preOrder func(*TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		preOrder(root.Left)
		preOrder(root.Right)
	}

	preOrder(root)
	return
}
