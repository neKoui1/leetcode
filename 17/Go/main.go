package main

import (
	"fmt"
	"strings"
)

func letterCombinations(digits string) (res []string) {
	if strings.Compare(digits, "") == 0 {
		return
	}
	n := len(digits)
	path := make([]byte, 0)
	value := []string{
		"abc",  //0->2
		"def",  //1->3
		"ghi",  //2
		"jkl",  //3
		"mno",  //4
		"pqrs", //5
		"tuv",  //6
		"wxyz", //7
	}

	var backtracking func(index int)
	backtracking = func(index int) {
		if len(path) == n {
			temp := string(path)
			res = append(res, temp)
			return
		}
		digit := int(digits[index] - '2')
		letter := value[digit]
		for i := 0; i < len(letter); i++ {
			path = append(path, letter[i])
			backtracking(index + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)

	return
}
func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
}
