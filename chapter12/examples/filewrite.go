package main

import (
	"os"
)

func main() {
	os.Stdout.WriteString("Hello World\n")
	f, _ := os.OpenFile("./chapter12/examples/test", os.O_CREATE|os.O_WRONLY, 0)
	defer f.Close()
	f.WriteString("hello, world in a file\n")
}
