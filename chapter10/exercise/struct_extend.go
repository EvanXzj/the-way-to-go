package main

import "fmt"

type A struct{ a int }
type B struct{ a, b int }

type C struct {
	A
	B
}

func main() {
	c := C{A{1}, B{2, 3}}
	fmt.Println(c)
	fmt.Println(c.A.a, c.b, c.B.a, c.B.b)
	// fmt.Println(c.a) 报错： ambiguous selector c.a
}
