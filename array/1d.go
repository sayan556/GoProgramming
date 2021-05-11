package main

import "fmt"

func main() {
	//Declare and Initialize an array
	var arr1 = [6]int{2, 3, 5, 7, 11, 13}
	for i := 0; i < 6; i++ {
		fmt.Print(arr1[i], " ")
	}
}
