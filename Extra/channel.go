package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
func main() {
	s1 := []int{1, 3, 5, 7, 9}
	s2 := []int{2, 4, 6, 8, 10}
	c := make(chan int)
	go sum(s1, c) //s[:len(s)/2 -> [0, 1]
	go sum(s2, c) //s[len(s)/2:] -> [2, 3, 4]
	x, y := <-c, <-c
	fmt.Printf("1st Sum = %d\n2nd Sum = %d\nTotal Sum = %d", x,
		y, x+y)
}
