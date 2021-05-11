package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Enter values : ")
	fmt.Scan(&a, &b)
	fmt.Println("Before swap : ")
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println("After swap")
	fmt.Println(a, b)
}
func swap(x *int, y *int) {
	var temp int
	temp = *x /* save the value at address x */
	*x = *y   /* put y into x */
	*y = temp /* put temp into y */
}
