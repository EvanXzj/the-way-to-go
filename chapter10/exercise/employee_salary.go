package main

import "fmt"

type employee struct {
	salary float32
}

func (e *employee) giveRaise(pct float32) {
	e.salary += e.salary * pct
}

func main() {
	/* create an employee instance */
	e := new(employee)
	e.salary = 100000
	/* call our method */
	e.giveRaise(0.04)
	fmt.Printf("Employee now makes %f", e.salary)
}
