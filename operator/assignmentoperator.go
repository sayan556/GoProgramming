package main

import "fmt"

func main() {
	a := 1
	var c int
	c = a
	fmt.Println(c)
	c += a
	fmt.Println(c)
	c -= a
	fmt.Println(c)
	c *= a
	fmt.Println(c)
	c /= a
	fmt.Println(c)
	c %= a
	fmt.Println(c)
	c <<= a
	fmt.Println(c)
	c >>= a
	fmt.Println(c)
	c &= a
	fmt.Println(c)
	c ^= a
	fmt.Println(c)
	c |= a
	fmt.Println(c)
}
