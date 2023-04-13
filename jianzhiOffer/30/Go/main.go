package main

import "math"

type MinStack struct {
	data []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	m := MinStack{
		[]int{},
		[]int{},
	}
	m.min = append(m.min, math.MaxInt)
	return m
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	n := len(this.min)
	if x < this.min[n-1] {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[n-1])
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return 0
	}
	return this.data[len(this.data)-1]
}

func (this *MinStack) Min() int {
	if len(this.min) == 1 || len(this.min) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
func main() {

}
