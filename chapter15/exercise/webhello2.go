package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// func helloServer(w http.ResponseWriter, req *http.Request) {
// 	fmt.Println("Inside helloServer hanndler")
// 	ps := strings.Split(req.URL.Path, "/")
// 	name := ps[2]
// 	if ps[1] == "shouthello" {
// 		name = strings.ToUpper(name)
// 	}
// 	fmt.Fprintf(w, "Hello,"+ name)
// }

func helloServer(w http.ResponseWriter, req *http.Request) {
	remPartOfURL := req.URL.Path[len("/hello/"):]
	fmt.Fprintf(w, "Hello %s!", remPartOfURL)
}

func shouthelloHandler(w http.ResponseWriter, req *http.Request) {
	remPartOfURL := req.URL.Path[len("/shouthello/"):] //get everything after the /shouthello/ part of the URL
	fmt.Fprintf(w, "Hello %s!", strings.ToUpper(remPartOfURL))
}

func main() {
	http.HandleFunc("/hello/", helloServer)
	http.HandleFunc("/shouthello/", shouthelloHandler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
