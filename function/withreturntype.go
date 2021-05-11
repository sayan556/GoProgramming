package main

import "fmt"

func sqr(x int) int {
	result := x * x
	return result
}

func main() {
	var num int
	fmt.Print("Enter number : ")
	fmt.Scan(&num)
	result := sqr(num)
	fmt.Println("Result = ", result)
}
