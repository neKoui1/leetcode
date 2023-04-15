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
	node2Pre := make(map[*TreeNode]*TreeNode, 0)
	node2Height := make(map[*TreeNode]int, 0)
	var dfs func(node *TreeNode, pre *TreeNode)
	dfs = func(node *TreeNode, pre *TreeNode) {
		if root == nil {
			return
		}
		if pre == nil {
			node2Height[node] = 1
			node2Pre[node] = nil
		} else {
			node2Height[node] = 1 + node2Height[pre]
			node2Pre[node] = pre
		}
		dfs(root.Left, root)
		dfs(root.Right, root)
	}
	dfs(root, nil)
	fmt.Println(node2Pre)
	fmt.Println(node2Height)
	return nil
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
