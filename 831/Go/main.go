package main

import (
	"fmt"
	"strings"
)

func maskPII(s string) (res string) {
	flagIsEmail := strings.Contains(s, "@")
	if flagIsEmail {
		s = strings.ToLower(s)
		strArr := strings.Split(s, "@")
		res = string(strArr[0][0]) + "*****" + string(strArr[0][len(strArr[0])-1])
		res += "@" + strArr[1]
	} else {
		s = strings.Replace(s, "(", "", -1)
		s = strings.Replace(s, ")", "", -1)
		s = strings.Replace(s, "-", "", -1)
		s = strings.Replace(s, "+", "", -1)
		s = strings.Replace(s, " ", "", -1)
		n := len(s)
		suffix := s[n-4:]
		switch n {
		case 10:
			res = "***-***-"
		case 11:
			res = "+*-***-***-"
		case 12:
			res = "+**-***-***-"
		case 13:
			res = "+***-***-***-"
		default:
			res = ""
		}
		res += suffix
	}
	return
}
func main() {
	fmt.Println(maskPII("LeetCode@LeetCode.com"))
	fmt.Println(maskPII("AB@qq.com"))
	fmt.Println(maskPII("1(234)567-890"))
	fmt.Println(maskPII("86-(10)12345678"))
	fmt.Println(maskPII("(3906)2 07143 711"))
}
