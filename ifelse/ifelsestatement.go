package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter age : ")
	fmt.Scanln(&age)

	/* check the boolean condition using if statement */
	if age > 18 {
		/* if condition is true then print the following */
		fmt.Println("You are adult")
	} else {
		fmt.Println("You are not adult")
	}
}
