package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	cnt := 0
	ans := 0
	var traverse func(cur *TreeNode)
	traverse = func(cur *TreeNode) {
		if cur.Right != nil {
			traverse(cur.Right)
		}
		cnt++
		if cnt == k {
			ans = cur.Val
			return
		}
		if cur.Left != nil {
			traverse(cur.Left)
		}
	}
	traverse(root)
	return ans
}
func main() {

}
