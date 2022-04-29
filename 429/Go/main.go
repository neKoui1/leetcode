package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	nodes := []*Node{root}
	for nodes != nil {
		level := make([]int, 0)
		tmp := nodes
		nodes = nil
		for _, v := range tmp {
			level = append(level, v.Val)
			nodes = append(nodes, v.Children...)
		}
		res = append(res, level)
	}
	return res
}
