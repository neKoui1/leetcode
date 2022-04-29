package main

import "fmt"

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

/*
	如果全为1， 则root.Val = true
	………………0， ………………………………false
	如果生成的是根节点，则root.IsLeaf = false,root.Val = true
*/
func construct(grid [][]int) *Node {
	if grid == nil {
		return nil
	}
	root := new(Node)
	root.IsLeaf = false
	root.Val = true

	m := len(grid)

	/*
		第一种情况，二维数组只有一个元素，此时只生成一个结点
	*/
	if m == 1 {
		root.IsLeaf = true
		if grid[0][0] == 0 {
			root.Val = false
			return root
		}
		return root
	}
	/*
		第二种情况，二维数组有四个元素，此时需要比较
		才能得出是生成一个结点还是生成四个结点
	*/
	if m == 2 {
		flag := true
		temp := grid[0][0]
		for _, v := range grid {
			if flag {
				for _, v2 := range v {
					if temp != v2 {
						flag = false
						break
					}
				}
			}
		}

		if flag {
			root.IsLeaf = true
			if grid[0][0] == 0 {
				root.Val = false
				return root
			}
			return root
		}

		root.TopLeft = construct([][]int{{grid[0][0]}})
		root.TopRight = construct([][]int{{grid[0][1]}})
		root.BottomLeft = construct([][]int{{grid[1][0]}})
		root.BottomRight = construct([][]int{{grid[1][1]}})

		return root
	}

	/*
		第3种情况
		在已知矩阵大小不等于1和4的情况下划分为4块
		然后为root的四叉赋值
	*/
	FLAG := true
	for _, v := range grid {
		if FLAG {
			temp := grid[0][0]
			for _, v2 := range v {
				if temp != v2 {
					FLAG = false
					break
				}
			}
		}
	}

	if FLAG {
		if grid[0][0] == 0 {
			return &Node{false, true, nil, nil, nil, nil}
		}
		return &Node{true, true, nil, nil, nil, nil}
	}

	flag1, flag2, flag3, flag4 := true, true, true, true
	res1 := make([][]int, 0)
	res2 := make([][]int, 0)
	res3 := make([][]int, 0)
	res4 := make([][]int, 0)
	for i := 0; i < m/2; i++ {
		line1 := make([]int, 0)
		line2 := make([]int, 0)
		line3 := make([]int, 0)
		line4 := make([]int, 0)
		for j := 0; j < m/2; j++ {
			line1 = append(line1, grid[i][j])
			line2 = append(line2, grid[i][j+m/2])
			line3 = append(line3, grid[i+m/2][j])
			line4 = append(line4, grid[i+m/2][j+m/2])
		}
		res1 = append(res1, line1)
		res2 = append(res2, line2)
		res3 = append(res3, line3)
		res4 = append(res4, line4)
	}

	for _, v := range res1 {
		if flag1 {
			temp := res1[0][0]
			for _, v2 := range v {
				if temp != v2 {
					flag1 = false
					break
				}
			}
		}
	}

	for _, v := range res2 {
		if flag2 {
			temp := res2[0][0]
			for _, v2 := range v {
				if temp != v2 {
					flag2 = false
					break
				}
			}
		}
	}

	for _, v := range res3 {
		if flag3 {
			temp := res3[0][0]
			for _, v2 := range v {
				if temp != v2 {
					flag3 = false
					break
				}
			}
		}
	}

	for _, v := range res4 {
		if flag4 {
			temp := res4[0][0]
			for _, v2 := range v {
				if temp != v2 {
					flag4 = false
					break
				}
			}
		}
	}

	if flag1 {
		if res1[0][0] == 0 {
			root.TopLeft = &Node{false, true, nil, nil, nil, nil}
		} else {
			root.TopLeft = &Node{true, true, nil, nil, nil, nil}
		}
	} else {
		root.TopLeft = construct(res1)
	}

	if flag2 {
		if res2[0][0] == 0 {
			root.TopRight = &Node{false, true, nil, nil, nil, nil}
		} else {
			root.TopRight = &Node{true, true, nil, nil, nil, nil}
		}
	} else {
		root.TopRight = construct(res2)
	}

	if flag3 {
		if res3[0][0] == 0 {
			root.BottomLeft = &Node{false, true, nil, nil, nil, nil}
		} else {
			root.BottomLeft = &Node{true, true, nil, nil, nil, nil}
		}
	} else {
		root.BottomLeft = construct(res3)
	}

	if flag4 {
		if res4[0][0] == 0 {
			root.BottomRight = &Node{false, true, nil, nil, nil, nil}
		} else {
			root.BottomRight = &Node{true, true, nil, nil, nil, nil}
		}
	} else {
		root.BottomRight = construct(res4)
	}

	return root
}

func main() {
	res := make([][]int, 0)
	fmt.Println(len(res))
}
