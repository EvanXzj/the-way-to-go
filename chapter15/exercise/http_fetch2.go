package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the url...")
	url, err := inputReader.ReadString('\n')
	url = strings.TrimSuffix(url, "\n")
	checkError(err)
	resp, err := http.Get(url)
	checkError(err)
	data, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	fmt.Printf("Got: %q", string(data))
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
