package main

import (
	"fmt"
	"strconv"
)

func calPoints(ops []string) int {
	score := make([]int, 0)
	for _, v := range ops {
		s, err := strconv.Atoi(v)
		if err != nil {
			switch v {
			case "+":
				score = append(score, score[len(score)-1]+score[len(score)-2])
			case "C":
				score = score[:len(score)-1]
			case "D":
				score = append(score, score[len(score)-1]*2)
			default:
				break
			}
		} else {
			score = append(score, s)
		}
	}
	sum := 0
	for _, v := range score {
		sum += v
	}

	return sum
}

func main() {
	ops := []string{"5", "2", "C", "D", "+"}
	fmt.Println(calPoints(ops))
	ops = []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	fmt.Println(calPoints(ops))
}
