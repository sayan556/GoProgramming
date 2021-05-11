package main

import "fmt"

//Global variable
var e float32 = 3.3

func main() {
	//Local variable
	var f, g, h = true, false, "hello!!"
	i := 4
	fmt.Println(e, f, g, h, i)
}
