// package main

// import "fmt"

// func Max(i, j int) int {
// 	if i > j {
// 		return i
// 	}
// 	return j
// }

// func lengthOfLIS(nums []int) int {
// 	n := len(nums)
// 	if n == 0 {
// 		return 0
// 	}
// 	dp := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		dp[i] = 1
// 		for j := 0; j < i; j++ {
// 			if nums[j] < nums[i] {
// 				dp[i] = Max(dp[i], dp[j]+1)
// 			}
// 		}
// 	}
// 	ans := 0
// 	for _, v := range dp {
// 		ans = Max(ans, v)
// 	}
// 	return ans
// }
// func main() {
// 	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
// 	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
// 	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
// 	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
// }

package main

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func lengthOfLIS(nums []int) int {
	ans := 0
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
	}
	for _, v := range dp {
		ans = Max(ans, v)
	}

	return ans
}

func main() {

}
