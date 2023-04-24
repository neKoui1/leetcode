package main

/**
给你一个字符串 s ，找出它的所有子串并按字典序排列，返回排在最后的那个子串。
*/
func lastSubstring(s string) string {
	n := len(s)
	i := 0
	for j, k := 1, 0; j+k < n; {
		if s[i+k] == s[j+k] {
			k++
		} else if s[i+k] < s[j+k] {
			i += k + 1
			k = 0
			if i >= j {
				j = i + 1
			}
		} else {
			j += k + 1
			k = 0
		}
	}
	return s[i:]
}
func main() {

}
