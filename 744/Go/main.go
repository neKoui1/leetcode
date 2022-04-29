package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	res := letters[0]
	for _, v := range letters {
		if v > target {
			res = v
			break
		}
	}
	return res
}
func main() {
	letters := []byte{'c', 'f', 'j'}
	fmt.Println(nextGreatestLetter(letters, 'a'))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c')))
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd'))
}
