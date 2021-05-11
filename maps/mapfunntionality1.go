package main

import "fmt"

func main() {
	sub := make(map[string]int)
	sub["math"] = 42
	fmt.Println("The value:", sub["math"])
	sub["math"] = 48
	fmt.Println("The value:", sub["math"])
	delete(sub, "math")
	fmt.Println("The value:", sub["math"])
	v, ok := sub["math"]
	fmt.Println("The value:", v, "Present?", ok)
}
