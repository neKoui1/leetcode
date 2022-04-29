package main

import "fmt"

func freqAlphabets(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if i+2 < len(s) && s[i+2] == '#' {
			res += string((s[i]-'0')*10 + s[i+1] - '1' + 'a')
			i += 2
		} else {
			res += string(s[i] - '1' + 'a')
		}
	}
	return res
}

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
	fmt.Println(freqAlphabets("1326#"))
}
