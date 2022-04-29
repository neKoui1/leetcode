package main

import "fmt"

func imageSmoother(img [][]int) [][]int {
	m := len(img)
	n := len(img[0])
	row := make([]int, 0)
	res := make([][]int, 0)
	if m == 1 && n == 1 {
		return img
	} else if m == 1 && n == 2 {
		temp := (img[0][0] + img[0][1]) / 2
		row = append(row, temp)
		row = append(row, temp)
		res = append(res, row)
	} else if m == 1 && n > 2 {
		for index, val := range img {
			for i := range val {
				if index == 0 {
					if i == 0 {
						temp := (img[0][0] + img[0][1]) / 2
						row = append(row, temp)
					} else if i == n-1 {
						temp := (img[0][i] + img[0][i-1]) / 2
						row = append(row, temp)
						res = append(res, row)
					} else {
						temp := (img[0][i-1] + img[0][i] + img[0][i+1]) / 3
						row = append(row, temp)
					}
				}
			}
		}
	} else if m == 2 && n == 1 {
		temp := (img[0][0] + img[1][0]) / 2
		row = append(row, temp)
		res = append(res, row)
		res = append(res, row)
	} else if m == 2 && n == 2 {
		temp := (img[0][0] + img[0][1] + img[1][0] + img[1][1]) / 4
		row = append(row, temp)
		row = append(row, temp)
		res = append(res, row)
		res = append(res, row)
	} else if m > 2 && n == 1 {
		for index := range img {
			if index == 0 {
				temp := (img[0][0] + img[1][0]) / 2
				row = append(row, temp)
				res = append(res, row)
				row = []int{}
			} else if index == m-1 {
				temp := (img[index][0] + img[index-1][0]) / 2
				row = append(row, temp)
				res = append(res, row)
				row = nil
			} else {
				temp := (img[index-1][0] + img[index][0] + img[index+1][0]) / 3
				row = append(row, temp)
				res = append(res, row)
				row = []int{}
			}
		}
	} else if m > 2 && n == 2 {
		for index := range img {
			if index == 0 {
				temp := (img[0][0] + img[0][1] + img[1][0] + img[1][1]) / 4
				row = append(row, temp)
				row = append(row, temp)
				res = append(res, row)
				row = []int{}
			} else if index == m-1 {
				temp := (img[index][0] + img[index][1] + img[index-1][0] + img[index-1][1]) / 4
				row = append(row, temp)
				row = append(row, temp)
				res = append(res, row)
				row = []int{}
			} else {
				temp := (img[index][0] + img[index][1] + img[index-1][0] + img[index-1][1] + img[index+1][0] + img[index+1][1]) / 6
				row = append(row, temp)
				row = append(row, temp)
				res = append(res, row)
				row = []int{}
			}
		}
	} else if m == 2 && n > 2 {
		for index, val := range img {
			for i := range val {
				if index == 0 {
					if i == 0 {
						temp := (img[0][0] + img[0][1] + img[1][0] + img[1][1]) / 4
						row = append(row, temp)
					} else if i == n-1 {
						temp := (img[0][i] + img[0][i-1] + img[1][i] + img[1][i-1]) / 4
						row = append(row, temp)
						res = append(res, row)
						row = []int{}
					} else {
						temp := (img[0][i] + img[0][i-1] + img[0][i+1] + img[1][i] + img[1][i-1] + img[1][i+1]) / 6
						row = append(row, temp)
					}
				} else {
					if i == 0 {
						temp := (img[0][0] + img[0][1] + img[1][0] + img[1][1]) / 4
						row = append(row, temp)
					} else if i == n-1 {
						temp := (img[0][i] + img[0][i-1] + img[1][i] + img[1][i-1]) / 4
						row = append(row, temp)
						res = append(res, row)
						row = []int{}
					} else {
						temp := (img[0][i] + img[0][i-1] + img[0][i+1] + img[1][i] + img[1][i-1] + img[1][i+1]) / 6
						row = append(row, temp)
					}
				}
			}
		}
	} else {
		for index, val := range img {
			for i := range val {
				if index == 0 {
					if i == 0 {
						temp := (img[0][0] + img[0][1] + img[1][0] + img[1][1]) / 4
						row = append(row, temp)
					} else if i == n-1 {
						temp := (img[0][n-1] + img[0][n-2] + img[1][n-1] + img[1][n-2]) / 4
						row = append(row, temp)
						res = append(res, row)
						row = []int{}
					} else {
						temp := (img[0][i] + img[0][i-1] + img[0][i+1] + img[1][i-1] + img[1][i] + img[1][i+1]) / 6
						row = append(row, temp)
					}
				} else if index == m-1 {
					if i == 0 {
						temp := (img[m-1][0] + img[m-1][1] + img[m-2][0] + img[m-2][1]) / 4
						row = append(row, temp)
					} else if i == n-1 {
						temp := (img[m-1][n-1] + img[m-1][n-2] + img[m-2][n-1] + img[m-2][n-2]) / 4
						row := append(row, temp)
						res = append(res, row)
						row = nil
					} else {
						temp := (img[m-1][i-1] + img[m-1][i] + img[m-1][i+1] + img[m-2][i-1] + img[m-2][i] + img[m-2][i+1]) / 6
						row = append(row, temp)
					}
				} else {
					if i == 0 {
						temp := (img[index][i] + img[index][i+1] + img[index-1][i] + img[index-1][i+1] + img[index+1][i] + img[index+1][i+1]) / 6
						row = append(row, temp)
					} else if i == n-1 {
						temp := (img[index][i-1] + img[index][i] + img[index-1][i-1] + img[index-1][i] + img[index+1][i-1] + img[index+1][i]) / 6
						row = append(row, temp)
						res = append(res, row)
						row = []int{}
					} else {
						temp := (img[index][i] + img[index][i+1] + img[index][i-1] + img[index-1][i] + img[index-1][i-1] + img[index-1][i+1] + img[index+1][i] + img[index+1][i-1] + img[index+1][i+1]) / 9
						row = append(row, temp)
					}
				}
			}
		}
	}

	return res
}

func main() {
	fmt.Println(imageSmoother([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	fmt.Println(imageSmoother([][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}))
	fmt.Println(imageSmoother([][]int{{1}}))
	fmt.Println(imageSmoother([][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}))
	fmt.Println(imageSmoother([][]int{{2, 5, 8}, {4, 0, 1}}))
	fmt.Println(imageSmoother([][]int{{2, 5, 8}, {4, 0, 1}}))
	fmt.Println(imageSmoother([][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10}}))
	fmt.Println(imageSmoother([][]int{{2, 5}, {8, 4}, {0, 1}}))
}
