package main

import "fmt"

func isRobotBounded(instructions string) bool {
	// 偏移量
	// vec[0] 向北 （向上）
	// vec[1] 向东 （右）
	// vec[2] 向南 （下）
	// vec[3] 向西 （左）
	direVec := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	direc := 0   // 初始方向
	x, y := 0, 0 // 初始坐标
	n := len(instructions)
	for i := 0; i < n; i++ {
		ins := instructions[i]
		if ins == 'G' {
			x += direVec[direc][0]
			y += direVec[direc][1]
		} else if ins == 'L' {
			direc += 3
			direc %= 4
		} else if ins == 'R' {
			direc++
			direc %= 4
		}
	}
	return direc != 0 || (x == 0 && y == 0)
}
func main() {
	fmt.Println(isRobotBounded("GGLLGG"))
	fmt.Println(isRobotBounded("GG"))
	fmt.Println(isRobotBounded("GL"))
	fmt.Println(isRobotBounded("GLRLGLLGLGRGLGL"))
}
