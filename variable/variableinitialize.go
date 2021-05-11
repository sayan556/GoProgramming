package main

import "fmt"

//Global variable
var c, d int = 1, 2
var e float32 = 3.3

func main() {
	//Local variable
	var f, g, h = true, false, "hello!!"
	fmt.Println(c, d, e, f, g, h)
}
