package main

import (
	"fmt"
	"sort"
)

/*
	单纯用哈希表交换键值对来排序的话会引起哈希冲突
	例如[1,2] hash[1]=1 hash[2]=1
	没有办法直接对val排序，交换后就变成了sw[1] = 1 sw[1] = 1
	所以自己写一个pair结构体出来比较好
*/
type Pair struct {
	Key int
	Cnt int
}

func topKFrequent(nums []int, k int) (res []int) {
	hash := make(map[int]int, 0)
	for _, v := range nums {
		hash[v]++
	}
	temp := make([]Pair, 0)
	for i, v := range hash {
		temp = append(temp, Pair{
			i,
			v,
		})
	}
	sort.SliceStable(temp, func(i, j int) bool {
		return temp[i].Cnt > temp[j].Cnt
	})
	for i := 0; i < k; i++ {
		res = append(res, temp[i].Key)
	}
	return
}
func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
	fmt.Println(topKFrequent([]int{1, 2}, 2))
}
