package main

import "fmt"

type MyStack struct {
	stack []int
}

func Constructor() MyStack {
	return MyStack{[]int{}}
}

func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyStack) Pop() int {
	m := len(this.stack) - 1
	res := this.stack[m]
	this.stack = this.stack[:m]
	return res
}

func (this *MyStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MyStack) Empty() bool {
	if len(this.stack) == 0 {
		return true
	} else {
		return false
	}
}
func main() {
	s := Constructor()
	fmt.Printf("s.Empty(): %v\n", s.Empty())
}
