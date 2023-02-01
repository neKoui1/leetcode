package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Key string
	Cnt int
}

func topKFrequent(words []string, k int) (res []string) {
	hash := make(map[string]int, 0)
	for _, v := range words {
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
		return temp[i].Key < temp[j].Key
	})
	sort.SliceStable(temp, func(i, j int) bool {
		return temp[i].Cnt > temp[j].Cnt
	})
	for i := 0; i < k; i++ {
		res = append(res, temp[i].Key)
	}
	return res
}
func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
	fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 3))
}
