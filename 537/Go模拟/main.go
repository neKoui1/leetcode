package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	使用IndexByte快速确定索引值
	去你妈的遍历
*/
func complexNumberMultiply(num1 string, num2 string) string {
	i := strings.IndexByte(num1, '+')
	j := strings.IndexByte(num2, '+')
	r1, _ := strconv.Atoi(num1[:i])
	c1, _ := strconv.Atoi(num1[i+1 : len(num1)-1])
	r2, _ := strconv.Atoi(num2[:j])
	c2, _ := strconv.Atoi(num2[j+1 : len(num2)-1])
	return fmt.Sprintf("%d+%di", r1*r2-c1*c2, r1*c2+r2*c1)
}

func main() {
	fmt.Println(complexNumberMultiply("78+-76i", "-86+72i"))
}
