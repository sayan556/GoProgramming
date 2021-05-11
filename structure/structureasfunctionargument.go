package main

import "fmt"

type Person struct {
	name  string
	age   int
	email string
	phone int
}

func main() {
	var p1, p2 Person
	p1.name = "Alex"
	p1.age = 30
	p1.email = "alex@test.com"
	p1.phone = 9876543210
	p2.name = "Bob"
	p2.age = 26
	p2.email = "bob@test.com"
	p2.phone = 1234567890
	printPerson(p1)
	printPerson(p2)
}
func printPerson(p Person) {
	fmt.Printf("Name : %s\tAge : %d\tEmail : %s\tPhone : %d\n", p.name, p.age, p.email,
		p.phone)
}
