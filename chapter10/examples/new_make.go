// 不能编译
package main

import "fmt"

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func main() {
	// OK
	y := new(Bar)
	(*y).thingOne = "Hello"
	(*y).thingTwo = 1

	// NOT OK
	z := make(Bar) // 编译错误： cannot make type Bar
	(*z).thingOne = "Hello"
	(*z).thingTwo = 1

	// OK
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	// NOT OK
	u := new(Foo)
	(*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
	(*u)["y"] = "world"
}
