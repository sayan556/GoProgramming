package main

import "fmt"

func percentage(c, s float64) {
	var amt float64
	var per float64
	if s > c {
		amt = s - c
		per = ((amt * 100) / c)
		fmt.Printf("profit = %f percentage", per)
	} else if c > s {
		amt = c - s
		per = ((amt * 100) / c)
		fmt.Printf("loss =%f percentage", per)
	} else {
		fmt.Println("No profit or loss")
	}
}
func main() {
	var a, b float64
	fmt.Println("Enter Cost price : ")
	fmt.Scanln(&a)
	fmt.Println("Enter Sell price : ")

	fmt.Scanln(&b)
	percentage(a, b)
}
