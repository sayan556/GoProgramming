package main

import "fmt"

type Rectangle interface {
	perimeter() float64
	area() float64
}
type Values struct {
	l float64
	b float64
}

func (v Values) perimeter() float64 {
	return 2 * (v.l + v.b)
}
func (v Values) area() float64 {
	return v.l * v.b
}
func main() {
	var t Rectangle
	t = Values{10, 14}
	fmt.Println("Perimeter of rectangle :", t.perimeter())
	fmt.Println("Area of rectangle :", t.area())
}
