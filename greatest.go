package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Enter 1st number ")
	fmt.Scanln(&a)
	fmt.Print("Enter 2nd number ")
	fmt.Scanln(&b)
	fmt.Print("Enter 3rd number ")
	fmt.Scanln(&c)

	if a > b && a > c {
		fmt.Println(a, "is greatest number")
	} else if b > a && b > c {
		fmt.Println(b, "is greatest number")
	} else {
		fmt.Println(c, "is greatest number")
	}

	if a < b && a < c {
		fmt.Println(a, "is smallest number")
	} else if b < a && b < c {
		fmt.Println(b, "is smallest number")
	} else {
		fmt.Println(c, "is smallest number")
	}
}

// package main

// import "fmt"

// func Largest(number1, number2, number3 int) {
// 	if number1 >= number2 && number1 >= number3 {
// 		fmt.Println("Largest of three numbers: ", number1)
// 	} else if number2 >= number1 && number2 >= number3 {
// 		fmt.Println("Largest of three numbers: ", number2)
// 	} else {
// 		fmt.Println("Largest of three numbers: ", number3)
// 	}
// }

// func Smallest(number1, number2, number3 int) {
// 	if number1 <= number2 && number1 <= number3 {

// 		fmt.Println("Smallest of three numbers: ", number1)
// 	} else if number2 <= number1 && number2 <= number3 {
// 		fmt.Println("Smallest of three numbers: ", number2)
// 	} else {
// 		fmt.Println("Smallest of three numbers: ", number3)
// 	}
// }

// func main() {
// 	var n1, n2, n3 int
// 	fmt.Print("Enter First Number:")
// 	fmt.Scanln(&n1)
// 	fmt.Print("Enter Second Number:")
// 	fmt.Scanln(&n2)
// 	fmt.Print("Enter Third Number:")
// 	fmt.Scanln(&n3)

// 	Smallest(n1, n2, n3)
// 	Largest(n1, n2, n3)
// }
