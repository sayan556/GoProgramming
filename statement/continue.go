package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i < 10 {
			continue
		} else {
			fmt.Println(i)
		}
	}
}
