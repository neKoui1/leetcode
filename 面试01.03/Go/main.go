package main

import "fmt"

func replaceSpaces(S string, length int) string {
	var res []byte
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			res = append(res, "%20"...)
		} else {
			res = append(res, S[i])
		}
	}

	return string(res)
}
func main() {
	fmt.Println(replaceSpaces("Mr John Smith    ", 13))
	fmt.Println(replaceSpaces("               ", 5))

}
