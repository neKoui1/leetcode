package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	var swapLeftRight func(node *TreeNode)
	swapLeftRight = func(node *TreeNode) {
		if node == nil {
			return
		}
		node.Left, node.Right = node.Right, node.Left
		swapLeftRight(node.Left)
		swapLeftRight(node.Right)
	}
	swapLeftRight(root)
	return root
}
func main() {

}
