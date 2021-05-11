package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hello"
	fmt.Println(str1)
	str2 := "World"
	str3 := str1 + str2
	fmt.Println(str3)
	greetings := []string{"Hello", "world!"}
	fmt.Println(strings.Join(greetings, " "))
}
