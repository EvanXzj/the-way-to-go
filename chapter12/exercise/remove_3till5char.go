package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("goprogram")
	outputFile, _ := os.Open("goprogramT", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	var outputString string
	for {
		// inputString, readerError := inputReader.ReadString('\n')
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			break
		}
		//fmt.Printf("The input was: --%s--", inputString)
		if len(inputString) < 3 {
			outputString = "\n"
		} else if len(inputString) < 5 {
			outputString = string([]byte(inputString)[2:len(inputString)]) + "\n"
		} else {
			outputString = string([]byte(inputString)[2:5]) + "\n"
		}
		//fmt.Printf("The output was: --%s--", outputString)
		_, err := outputWriter.WriteString(outputString)
		//fmt.Printf("Number of bytes written %d\n", n)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	outputWriter.Flush() // write content to file
	fmt.Println("Conversion done")
}
