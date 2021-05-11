package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Create("create.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File created")
	}
}
