package main

/*
	抄的
*/
func findDuplicates(nums []int) (res []int) {
	for _, v := range nums {
		if v > 0 {
			v = -v
		}
		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		} else {
			res = append(res, v)
		}
	}
	return
}
func main() {

}
