package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	var helper func(n1, n2 *TreeNode) bool
	helper = func(n1, n2 *TreeNode) bool {
		if n2 == nil {
			return true
		}
		if n1 == nil {
			return false
		}
		if n1.Val == n2.Val {
			return helper(n1.Left, n2.Left) && helper(n1.Right, n2.Right)
		} else {
			return false
		}
	}

	if A.Val == B.Val {
		if helper(A, B) {
			return true
		}
	}

	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
func main() {

}
