package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	i int
}

func (s *Simple) Get() int {
	return s.i
}

func (s *Simple) Set(u int) {
	s.i = u
}

func fI(it Simpler) int {
	it.Set(5)
	return it.Get()
}

func main() {
	var s Simple
	fmt.Println(fI(&s))
}
