package main

import "fmt"

func swap(x, y int) (int, int) {
	var temp int
	temp = x /* save the value of x */
	x = y    /* put y into x */
	y = temp /* put temp into y */
	return x, y
}
func main() {
	var a, b int
	fmt.Println("Enter values : ")
	fmt.Scan(&a, &b)
	fmt.Println("Before swap : ")
	fmt.Println(a, b)
	a, b = swap(a, b)
	fmt.Println("After swap")
	fmt.Println(a, b)
}
