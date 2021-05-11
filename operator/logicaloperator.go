package main

import "fmt"

func main() {
	a, b, c, d := 10, 20, 15, true
	fmt.Println((a > b) && (b > c))
	fmt.Println((a > b) || (b > c))
	fmt.Println(!d)

	num := 10
	fmt.Println("Number :", num)
	fmt.Printf("Number : %v\n", num)
}
