package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer hanndler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:]) // 去掉前面的斜杠
}

func main() {
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":9527", nil))
}
