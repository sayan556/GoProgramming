package main

import "fmt"

func main() {
	/* local variable definition */
	var b string
	var a int
	fmt.Println("Press 1 for BENGALI, 2 for HINDI, 3 for ENGLISH ")
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&a)
	switch a {
	case 1:
		b = "BENGALI"
	case 2:
		b = "HINDI"
	case 3:
		b = "ENGLISH"
	default:
		b = " not correct"
	}
	fmt.Printf("Your choice is%s\n", b)
}
