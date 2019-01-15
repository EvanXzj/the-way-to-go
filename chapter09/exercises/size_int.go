package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int = 10
	size := unsafe.Sizeof(i)
	fmt.Println("The size of an int is:", size) // 32位机器： 4， 64位机器：8
}
