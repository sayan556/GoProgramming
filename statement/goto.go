package main

import "fmt"

func main() {
	var age int
election:
	fmt.Print("Enter your age : ")
	fmt.Scan(&age)
	if age < 18 {
		fmt.Println("You are not eligible to vote")
		goto election
	} else {
		fmt.Println("You are eligible to vote")
	}
}
