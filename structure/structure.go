package main

import "fmt"

type Info struct {
	name string
	age  int
}

func main() {
	fmt.Println(Info{"Alex", 23})
}
