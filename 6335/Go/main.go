package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Val = 0
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		temp := make([]*TreeNode, len(queue))
		copy(temp, queue)
		queue = nil
		nextLevelSum := 0
		for _, v := range temp {
			if v.Left != nil {
				queue = append(queue, v.Left)
				nextLevelSum += v.Left.Val
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
				nextLevelSum += v.Right.Val
			}
		}
		for _, v := range temp {
			nextLevelChildSum := 0
			if v.Left != nil {
				nextLevelChildSum += v.Left.Val
			}
			if v.Right != nil {
				nextLevelChildSum += v.Right.Val
			}
			if v.Left != nil {
				v.Left.Val = nextLevelSum - nextLevelChildSum
			}
			if v.Right != nil {
				v.Right.Val = nextLevelSum - nextLevelChildSum
			}
		}
	}

	return root
}
func main() {
	root := new(TreeNode)
	root.Val = 5
	first := new(TreeNode)
	second := new(TreeNode)
	first.Val = 4
	second.Val = 9
	root.Left = first
	root.Right = second
	ff := new(TreeNode)
	fs := new(TreeNode)
	ff.Val = 1
	fs.Val = 10
	first.Left = ff
	first.Right = fs
	ss := new(TreeNode)
	ss.Val = 7
	second.Right = ss
	fmt.Println(replaceValueInTree(root))
}
