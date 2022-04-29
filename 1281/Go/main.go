package main

import "fmt"

func subtractProductAndSum(n int) int {
	list := make([]int, 0)
	for n != 0 {
		list = append(list, n%10)
		n /= 10
	}
	sum := 0
	mult := 1
	for _, v := range list {
		sum += v
		mult *= v
	}

	return (mult - sum)
}

func main() {
	fmt.Println(subtractProductAndSum(234))
	fmt.Println(subtractProductAndSum(4421))
}
