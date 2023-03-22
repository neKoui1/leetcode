package main

import "fmt"

func convertTemperature(celsius float64) (res []float64) {
	res = append(res, celsius+273.15)
	res = append(res, celsius*1.80+32.00)
	return res
}
func main() {
	fmt.Println(convertTemperature(36.50))
	fmt.Println(convertTemperature(122.11))
}
