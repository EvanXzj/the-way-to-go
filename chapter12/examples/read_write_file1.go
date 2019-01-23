package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inf := "./chapter12/examples/products.txt"
	ouf := "./chapter12/examples/products_copy.txt"
	buf, err := ioutil.ReadFile(inf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read File Error: %s\n", err)
	}
	fmt.Printf("File Content: %s\n", string(buf))
	err = ioutil.WriteFile(ouf, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}
