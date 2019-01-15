package main

import (
	"fmt"
	"runtime"
)

func main() {
	f := fbi(10)

	ret := f()

	fmt.Println(ret)
}

func fbi(n int) func() int{
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d", file, line)
	}
	where()
	a, b := 1, 1
	return func() int {
		for i:=0; i < n; i++ {
			a, b = b, a+b
		}

		return a
	}
}