package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("Enter name : ")
	//take user input and add it to name
	fmt.Scanln(&name)
	fmt.Print("Enter age : ")
	//take user input and add it to name
	fmt.Scanln(&age)
	fmt.Println("Your name : ", name, " Age : ", age)
}
