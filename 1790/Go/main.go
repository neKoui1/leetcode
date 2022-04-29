package main

import "fmt"

//构建一个哈希映射
func areAlmostEqual(s1 string, s2 string) bool {
	hash := make(map[int]rune)
	cnt := 0
	index := make([]int, 0)
	val := make([]rune, 0)
	for i, v := range s1 {
		hash[i] = v
	}
	for i, v := range s2 {
		if hash[i] != v {
			index = append(index, i)
			val = append(val, v)
			cnt++
		}
	}
	if cnt == 0 {
		return true
	} else if cnt == 2 {
		if hash[index[1]] == val[0] && hash[index[0]] == val[1] {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func main() {
	fmt.Println(areAlmostEqual("bank", "kanb"))
	fmt.Println(areAlmostEqual("attack", "defend"))
	fmt.Println(areAlmostEqual("kelb", "kelb"))
	fmt.Println(areAlmostEqual("abcd", "dcba"))
	fmt.Println(areAlmostEqual("caa", "aaz"))
}
