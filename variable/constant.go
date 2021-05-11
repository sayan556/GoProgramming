package main

import "fmt"

//Global constant variable
const Pi = 3.14

func main() {

	//Local constant variable
	const c = "World"
	fmt.Println("Hello", c)
	fmt.Println("Happy", Pi, "Day")
}
