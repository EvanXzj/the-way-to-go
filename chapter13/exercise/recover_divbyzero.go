package main

import "fmt"

func badCall() {
	a, b := 10, 0
	n := a / b
	fmt.Println(n)
}

func test() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panicing %s \n", r)
		}
	}()
	badCall()
	fmt.Printf("After bad call\n")
}

func main() {
	fmt.Printf("Calling test\n")
	test()
	fmt.Printf("Test completed\n")
}
