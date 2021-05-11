package main

import "fmt"

func main() {
	var a int
	fmt.Print("Enter a number : ")
	fmt.Scanln(&a)
	abc(a)
}
func abc(n int) {

	if n%2 == 0 {

		fmt.Println(n, "is Even number")
	} else {
		fmt.Println(n, "is Odd number")
	}
}
