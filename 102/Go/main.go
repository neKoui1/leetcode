package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {

		temp := []int{}
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}
func main() {

}
