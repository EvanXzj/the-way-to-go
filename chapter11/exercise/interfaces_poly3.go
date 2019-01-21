package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Shape struct{}

func (s Shape) Area() float32 {
	return -1
}

type Square struct {
	side float32
	Shape
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
	Shape
}

func (r *Rectangle) Area() float32 {
	return r.length * r.width
}

type Circle struct {
	radius float32
	Shape
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func main() {
	s := Shape{}
	q := &Square{5, s}
	r := &Rectangle{5, 3, s}
	c := &Circle{2.5, s}

	shapes := []Shaper{q, r, c, s}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
