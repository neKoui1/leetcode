package main

import (
	"fmt"
	"strconv"
)

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	res := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return res
	} else if root.Left != nil && root.Right == nil {
		res += "(" + tree2str(root.Left) + ")"
	} else if root.Left == nil && root.Right != nil {
		res += "()(" + tree2str(root.Right) + ")"
	} else {
		res += "(" + tree2str(root.Left) + ")"
		res += "(" + tree2str(root.Right) + ")"
	}
	return res
}

func main() {
	a := new(TreeNode)
	b := new(TreeNode)
	c := new(TreeNode)
	d := new(TreeNode)
	a.Val = 1
	b.Val = 2
	c.Val = 3
	d.Val = 4
	a.Left = b
	a.Right = c
	b.Right = d
	fmt.Println(tree2str(a))
}
