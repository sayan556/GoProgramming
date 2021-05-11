package main

import "fmt"

func main() {
	var str string = "Hello World"
	fmt.Println(str)
	fmt.Printf("%s\n", str)

	fmt.Println("Length : ", len(str))
}
