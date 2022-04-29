//罗马数字转整数

/*
str -> int ?

HashMap ??
*/

package main

import "fmt"

func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0

	for i, _ := range s {
		if i == len(s)-1 {
			sum += roman[s[i]]
			break
		}
		if roman[s[i]] >= roman[s[i+1]] {
			sum += roman[s[i]]
		} else {
			sum -= roman[s[i]]
		}
	}
	return sum
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	s := "1"
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", s[0])
	for _, v := range s {
		fmt.Printf("%T", v)
	}
}
