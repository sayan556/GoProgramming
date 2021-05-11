package main

import (
	"fmt"
)

type I interface {
	Area()
}
type Rectangle struct {
	l float64
	b float64
}

func (r Rectangle) Area() {
	fmt.Println("The Area of rectangle : ", r.l*r.b)
}

type Cone struct {
	r float64
	h float64
}

func (c Cone) Area() {
	fmt.Println("The Area of cone : ", 3.14*c.r*c.r*c.h/3)
}
func main() {
	var i I = Rectangle{10, 20}
	i.Area()
	i = Cone{10, 20}
	i.Area()
}
