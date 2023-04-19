package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var helper func(node1, node2 *TreeNode) bool
	helper = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil {
			return false
		}
		return node1.Val == node2.Val && helper(node1.Left, node2.Right) && helper(node1.Right, node2.Left)
	}

	return helper(root, root)
}
func main() {

}
