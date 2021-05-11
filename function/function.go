package main

import "fmt"

func checkValue(n1, n2 int) {
	if n1 > n2 {
		fmt.Println(n1, " is greater than ", n2)
	} else if n2 > n1 {
		fmt.Println(n2, " is greater than ", n1)
	} else {
		fmt.Println(n1, " is equal to ", n2)
	}
}
func main() {
	var num1, num2 int
	fmt.Println("Enter two number : ")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	checkValue(num1, num2)
}
