package main

import "fmt"

type Person struct {
	int
	string
	height float32
}

func main() {
	p := Person{23, "Evan", 1.75}
	fmt.Println(p.int, p.string, p.height)
	fmt.Println(p)
}
