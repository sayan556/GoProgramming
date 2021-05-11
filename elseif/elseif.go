package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter age : ")
	fmt.Scanln(&age)

	/* check the boolean condition using if statement */
	if age > 18 && age < 60 {
		fmt.Println("You are adult")
	} else if age >= 60 {
		fmt.Println("Senior citizen")
	} else {
		fmt.Println("You are not adult")
	}
}
