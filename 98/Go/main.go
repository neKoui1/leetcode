package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	arr := make([]int, 0)
	var spread func(root *TreeNode)
	spread = func(root *TreeNode) {
		if root == nil {
			return
		}
		spread(root.Left)
		arr = append(arr, root.Val)
		spread(root.Right)
	}
	spread(root)
	flag := true
	for i := 0; i < len(arr)-1 && flag; i++ {
		//如果有等于号，则依然return false
		if arr[i] >= arr[i+1] {
			flag = false
		}
	}
	return flag
}
func main() {

}
