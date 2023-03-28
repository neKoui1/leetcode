package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) (res [][]int) {

	path := make([]int, 0)
	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum -= root.Val
		path = append(path, root.Val)
		if sum == 0 && root.Left == nil && root.Right == nil {
			res = append(res, append([]int{}, path...))
		}
		dfs(root.Left, sum)
		dfs(root.Right, sum)
		path = path[:len(path)-1]
	}
	dfs(root, targetSum)
	return
}
func main() {

}
