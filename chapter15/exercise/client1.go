package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 打开链接
	conn, err := net.Dial("tcp", ":9527")
	checkError(err)

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedName := strings.Trim(clientName, "\n")
	// 给服务器发送信息直到程序退出：
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")
		if trimmedInput == "Q" {
			return
		}
		_, err := conn.Write([]byte(trimmedName + " says: " + trimmedInput))
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		panic("Error: " + err.Error())
	}
}
