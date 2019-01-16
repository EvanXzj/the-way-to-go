package main

import (
	"fmt"
	"./fibonacci"
)

func main() {
	n := 5
	op := "+"
	res := fibonacci.Fibonacci(op, n)
	fmt.Printf("op=%s, n=%d, res=%d\n", op, n, res)
	op = "*"
	res = fibonacci.Fibonacci(op, n)
	fmt.Printf("op=%s, n=%d, res=%d\n", op, n, res)
}
