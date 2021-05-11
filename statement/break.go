package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i < 5 {
			fmt.Println(i)
		} else {
			break
		}
	}
}
