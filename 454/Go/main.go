package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	hash := make(map[int]int)
	cnt := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			hash[nums1[i]+nums2[j]]++
		}
	}
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			if val, ok := hash[-nums3[i]-nums4[j]]; ok {
				cnt += val
			}
		}
	}
	return cnt
}
func main() {

}
