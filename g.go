package main

import "fmt"

func main() {
	b:= []byte{'g', 'o', 'l', 'a', 'n', 'g'}

	s1 := b[1:1]
	fmt.Println(string(s1), len(s1))

	slice1 := make([]int, 10)
    // load the array/slice:
    for i := 0; i < len(slice1); i++ {
        slice1[i] = 5 * i
    }

    // print the slice:
    for i := 0; i < len(slice1); i++ {
        fmt.Printf("Slice at %d is %d\n", i, slice1[i])
    }
    fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
    fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}