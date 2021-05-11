package main

import "fmt"

type Person struct {
	name  string
	age   int
	email string
	phone int
}

func main() {
	p := Person{"Alex", 28, "alex@test.com", 9876543210}
	fmt.Println("Name : ", p.name)
	fmt.Println("Age : ", p.age)
	fmt.Println("Email : ", p.email)
	fmt.Println("Phone : ", p.phone)
}
