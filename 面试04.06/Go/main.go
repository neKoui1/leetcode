package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	//var traverse func(root *TreeNode)
	//arr := make([]int, 0)
	//
	//traverse = func(root *TreeNode) {
	//	if root == nil {
	//		return
	//	}
	//	traverse(root.Left)
	//	arr = append(arr, root.Val)
	//	traverse(root.Right)
	//}
	//traverse(root)
	if root == nil {
		return nil
	}
	if p.Val >= root.Val {
		return inorderSuccessor(root.Right, p)
	}

	node := inorderSuccessor(root.Left, p)
	if node == nil {
		return root
	}
	return node
}
func main() {

}
