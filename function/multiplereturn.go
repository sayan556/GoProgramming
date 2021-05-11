package main

import "fmt"

func operation(a, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

func main() {
	var x, y int
	fmt.Println("Enter numbers : ")
	fmt.Scan(&x, &y)
	sum, sub := operation(x, y)
	fmt.Println("Summation = ", sum, " Subtraction = ", sub)
}
