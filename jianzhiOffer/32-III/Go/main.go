package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	cnt := 0
	for len(queue) != 0 {

		length := len(queue)
		temp := []int{}
		cnt++
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
		if cnt&1 != 1 {
			Reverse(temp)
		}
		res = append(res, temp)
	}
	return res
}
func main() {

}
