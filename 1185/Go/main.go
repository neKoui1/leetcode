package main

import "fmt"

func dayOfTheWeek(day int, month int, year int) (res string) {
	m := month
	if month == 1 || month == 2 {
		m += 12
		year--
	}
	c := year / 100
	y := year % 100
	d := day

	w := y + y/4 + c/4 - 2*c + (26 * (m + 1) / 10) + d - 1
	w = (w%7 + 7) % 7
	switch w {
	case 0:
		res = "Sunday"
	case 1:
		res = "Monday"
	case 2:
		res = "Tuesday"
	case 3:
		res = "Wednesday"
	case 4:
		res = "Thursday"
	case 5:
		res = "Friday"
	case 6:
		res = "Saturday"
	default:
		break
	}
	return
}
func main() {
	fmt.Println(dayOfTheWeek(25, 5, 2022))
	fmt.Println(dayOfTheWeek(29, 2, 2016))
	fmt.Println(dayOfTheWeek(29, 2, 2000))
	fmt.Println(dayOfTheWeek(7, 3, 2003))
}
