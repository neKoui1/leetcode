package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	nums []int
	set  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		[]int{},
		//map[val]=len-1
		//记录元素所在的数组的下标，因为每一次都是append在最后，所以就不引入新的变量直接算len-1
		map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.set[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	id, ok := this.set[val]
	if !ok {
		return false
	}
	last := len(this.nums) - 1
	this.nums[id] = this.nums[last]
	this.set[this.nums[id]] = id
	this.nums = this.nums[:last]
	delete(this.set, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
func main() {
	set := Constructor()
	set.Insert(1)
	fmt.Println(set.Remove(2))
	set.Insert(2)
	fmt.Println(set.GetRandom())
	set.Remove(1)
	fmt.Println(set.Insert(2))
	fmt.Println(set.GetRandom())
	fmt.Println()

	set = Constructor()
	fmt.Println(set.Insert(1))
	fmt.Println(set.Remove(2))
	fmt.Println(set.Insert(2))
	fmt.Println(set.GetRandom())
	fmt.Println(set.Remove(1))
	fmt.Println(set.Insert(2))
	fmt.Println(set.GetRandom())
	fmt.Println()

	s := make(map[int]struct{})
	fmt.Println(len(s))
}
