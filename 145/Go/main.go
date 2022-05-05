package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) (res []int) {
	var postOrder func(*TreeNode)
	postOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		postOrder(root.Left)
		postOrder(root.Right)
		res = append(res, root.Val)
	}
	postOrder(root)
	return
}
