package main

func sortColors(nums []int) {
	red, white, blue := 0, 0, 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == 2 {
			nums[blue] = 2
			blue++
		} else if nums[i] == 1 {
			nums[blue] = 2
			nums[white] = 1
			blue++
			white++
		} else {
			nums[blue] = 2
			nums[white] = 1
			nums[red] = 0
			blue++
			white++
			red++
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
}
