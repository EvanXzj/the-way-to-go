package main

import "fmt"

type Rectangle struct {
	length, width int
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func main() {
	r := Rectangle{4, 3}
	fmt.Println("Rectangle is: ", r)
	fmt.Println("Rectangle area is: ", r.Area())
	fmt.Println("Rectangle perimeter is: ", r.Perimeter())
}
