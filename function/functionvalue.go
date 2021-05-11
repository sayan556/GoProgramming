package main

import (
	"fmt"
)

func main() {
	cube := func(x int) int {
		return x * x * x
	}
	result := cube(5)
	fmt.Println("Result = ", result)
}
