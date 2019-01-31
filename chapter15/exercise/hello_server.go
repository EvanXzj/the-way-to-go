package main

import (
	"fmt"
	"net/http"
)

type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello")
}

func main() {
	var h hello
	http.ListenAndServe(":9000", h)
}
