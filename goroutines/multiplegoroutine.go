package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'j'; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%c ", i)
	}
}
func main() {
	go numbers()
	go alphabets()
	time.Sleep(5 * time.Second)
	fmt.Printf("\nmain terminated")
}
