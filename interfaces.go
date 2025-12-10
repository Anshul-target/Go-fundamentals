package main

import "fmt"

type Shape interface {
	Area() int
}

type Rectangle struct {
	length  int
	breadth int
}

type Circle struct {
	radius int
}

func (r Rectangle) Area() int {
	return r.breadth * r.length
}
func (c Circle) Area() int {
	return c.radius * c.radius
}

func calculateArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	// s := Circle{
	// 	radius: 12,
	// }
	var s Shape= Rectangle{
		length: 12,
		breadth: 13,
	}
	calculateArea(s)
}