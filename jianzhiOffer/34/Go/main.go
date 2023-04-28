package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	path := make([]int, 0)
	var backTracking func(cur *TreeNode, totalSum int)
	backTracking = func(cur *TreeNode, totalSum int) {
		if cur == nil {
			return
		}
		path = append(path, cur.Val)
		if totalSum+cur.Val == target && cur.Left == nil && cur.Right == nil {
			ans = append(ans, append([]int{}, path...))
			return
		}
		// 可能有负数相加的情况，所以这样剪枝会出错
		//if totalSum+cur.Val > target {
		//	return
		//}

		if cur.Left != nil {
			backTracking(cur.Left, totalSum+cur.Val)
			path = path[:len(path)-1]
		}
		if cur.Right != nil {
			backTracking(cur.Right, totalSum+cur.Val)
			path = path[:len(path)-1]
		}
	}
	backTracking(root, 0)
	return ans
}
func main() {

}
