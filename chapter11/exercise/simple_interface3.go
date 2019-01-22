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

type RSimple struct {
	i, j int
}

func (r *RSimple) Get() int {
	return r.j
}

func (r *RSimple) Set(u int) {
	r.j = u
}

func fI(it Simpler) int {
	switch it.(type) {
	case *Simple:
		it.Set(5)
		return it.Get()
	case *RSimple:
		it.Set(50)
		return it.Get()
	default:
		return 99
	}
}

func gI(any interface{}) int {
	if v, ok := any.(Simpler); ok {
		return v.Get()
	}
	return 0
}

func main() {
	var s Simple = Simple{6}
	fmt.Println(gI(&s)) // &s is required because Get() is defined with a receiver type pointer
	var r RSimple = RSimple{60, 60}
	fmt.Println(gI(&r))
}
