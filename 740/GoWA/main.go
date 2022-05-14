package main

import "fmt"

/*
	动态规划题
	一开始我没有看出来，还是思考的太简单了，需要的是“结果”最大值而非“每一次”的最大值
	所以如果使用暴力，应该穷举删除后的每一种结果，如果使用递归极其容易栈溢出
	所以此题马上看看别人的讲解再尝试自己写出来
	dp的题目前写不出来也算正常，可以理解
	也并不妨碍我骂自己废物
*/
func deleteAndEarn(nums []int) int {
	hash := make(map[int]int)
	for _, v := range nums {
		hash[v]++
	}
	max := 0
	record := 0
	for i, v := range hash {
		if i*v > max {
			max = i * v
			record = i
		}
	}
	delete(hash, record)
	if _, ok := hash[record-1]; ok {
		delete(hash, record-1)
	}
	if _, ok := hash[record+1]; ok {
		delete(hash, record+1)
	}
	cur := make([]int, 0)
	for i, v := range hash {
		tmp := v
		for tmp != 0 {
			cur = append(cur, i)
			tmp--
		}
	}
	if len(cur) == 1 {
		return max + cur[0]
	} else if len(cur) == 0 {
		return max
	} else {
		return max + deleteAndEarn(cur)
	}
}

func main() {
	fmt.Println(deleteAndEarn([]int{3, 4, 2}))
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))
	fmt.Println(deleteAndEarn([]int{8, 3, 4, 7, 6, 6, 9, 2, 5, 8, 2, 4, 9, 5, 9, 1, 5, 7, 1, 4}))
}
