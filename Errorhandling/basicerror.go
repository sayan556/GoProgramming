package main

import (
	"errors"
	"fmt"
)

func main() {
	var n1, n2 float64
	fmt.Print("Enter 1st number : ")
	fmt.Scanln(&n1)
	fmt.Print("Enter 2nd number : ")
	fmt.Scanln(&n2)
	result, error := division(n1, n2)
	if error != nil {
		fmt.Println("Result = ", result, error)
	} else {
		fmt.Println("Result = ", result)
	}
}
func division(x, y float64) (float64, error) {
	if x <= 0 || y <= 0 {
		return 0, errors.New("Math error : Zero is found from input value")
	} else {
		return (x / y), nil
	}
}
