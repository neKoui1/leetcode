package main

type RecentCounter struct {
	Num []int
}

func Constructor() RecentCounter {
	return RecentCounter{[]int{}}
}

func (this *RecentCounter) Ping(t int) int {
	this.Num = append(this.Num, t)
	cnt := 0
	for i := len(this.Num) - 1; i >= 0; i-- {
		if this.Num[i] < t-3000 {
			cnt = i + 1
			break
		}
	}
	return len(this.Num) - cnt
}
func main() {

}
