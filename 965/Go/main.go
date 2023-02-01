package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	也可以随便什么序遍历一遍
	用哈希表作为存储
	最后计算哈希表的len
	如果 len > 1
	则 return false
*/
func isUnivalTree(root *TreeNode) bool {
	flag := true
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			if root.Left.Val != root.Val {
				flag = false
				return
			} else {
				traverse(root.Left)
			}
		}
		if root.Right != nil {
			if root.Right.Val != root.Val {
				flag = false
				return
			} else {
				traverse(root.Right)
			}
		}

	}
	traverse(root)

	return flag
}

func main() {

}
