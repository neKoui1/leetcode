package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return nil
	}
	res = append(res, root.Val)
	for _, v := range root.Children {
		res = append(res, preorder(v)...)
	}
	return res
}
