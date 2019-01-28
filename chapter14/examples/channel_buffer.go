package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)

	go func() {
		time.Sleep(2 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()

	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
}

// Output:
// sending 10
// sent 10   // prints immediately
// no further output, because main() then stops
