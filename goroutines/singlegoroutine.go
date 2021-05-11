package main

import (
	"fmt"
	"time"
)

func hello(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}
func helloGo(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}
func main() {
	hello("Test")
	go helloGo("Testing")
	helloGo("Goroutines")
}
