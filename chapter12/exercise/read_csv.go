package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	title    string
	price    float64
	quantity int
}

func main() {
	bks := make([]Book, 1)
	file, err := os.Open("./chapter12/exercise/products.txt")
	if err != nil {
		log.Fatalf("Error %s opening file products.txt")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		hasSuffix := strings.HasSuffix(line, "\r\n")
		n := 1
		if hasSuffix {
			n = 2
		}
		line = string(line[:len(line)-n])

		strSlice := strings.Split(line, ";")
		book := new(Book)
		book.title = strSlice[0]
		book.price, err = strconv.ParseFloat(strSlice[1], 32)
		if err != nil {
			fmt.Printf("Error in file: %v\n", err)
		}
		book.quantity, err = strconv.Atoi(strSlice[2])
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}

		if bks[0].title == "" {
			bks[0] = *book
		} else {
			bks = append(bks, *book)
		}
	}

	fmt.Println("We have read the following books from the file: ")
	for _, bk := range bks {
		fmt.Println(bk)
	}
}
