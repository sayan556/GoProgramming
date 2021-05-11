package main

import "fmt"

func main() {
	a, b, c, d := 0, 1, 7, 3
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(c << d)
	fmt.Println(c >> d)
}
