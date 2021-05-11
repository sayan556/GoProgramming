package main

import "fmt"

func main() {
	a, b := 10, 2
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a += 1
	fmt.Println(a)
	b -= 1
	fmt.Println(b)
}
