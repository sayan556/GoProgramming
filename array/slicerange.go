package main

import "fmt"

var pow = []int{5, 10, 15, 20, 25, 30, 35, 40}

func main() {
	for i, v := range pow {
		fmt.Printf("Index = %d Value = %d\n", i, v)
	}
}
