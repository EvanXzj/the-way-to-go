package main

import "fmt"

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panicing %s\n", err)
		}
	}()

	badCall()
	fmt.Printf("After bad call\r\n")
}

func main() {
	fmt.Println("Calling test")
	test()
	fmt.Println("Test completed")
}
