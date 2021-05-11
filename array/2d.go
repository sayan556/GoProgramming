package main

import "fmt"

func main() {
	/* an array with 5 rows and 2 columns*/
	var a = [2][2]int{{0, 1}, {2, 3}}
	var i, j int
	/* output each array element's value */
	for i = 0; i < 2; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
	fmt.Println("Matrix form : ")
	for i = 0; i < 2; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("%d\t", a[i][j])
		}
		fmt.Println()
	}
}
