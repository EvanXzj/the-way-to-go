package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// Map of the clients: contains: clientname - 1 (active) / 0 - (inactive)
var mapUsers map[string]int

func main() {
	fmt.Println("Starting the server ...")
	mapUsers = make(map[string]int)
	// 创建listener
	listener, err := net.Listen("tcp", ":9527")
	checkError(err)
	// 监听并接受来自客户端的链接
	for {
		conn, err := listener.Accept()
		checkError(err)
		fmt.Println("got a connection")
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		length, err := conn.Read(buf)
		checkError(err)
		input := string(buf)
		if strings.Contains(input, ": SH") {
			fmt.Println("Server shutting down.")
			os.Exit(0)
		}
		if strings.Contains(input, ": WHO") {
			DisplayList()
		}
		ix := strings.Index(input, "says")
		clName := input[0 : ix-1]
		mapUsers[clName] = 1
		fmt.Printf("Received data: %v\n", string(buf[:length]))
	}
}

// advantage: code is cleaner,
// disadvantage:  the server process has to stop at any error:
//                a simple return continues in the function where we came from!
func checkError(err error) {
	if err != nil {
		panic("Error: " + err.Error())
	}
}

func DisplayList() {
	fmt.Println("--------------------------------------------")
	fmt.Println("This is the client list: 1=active, 0=inactive")
	for key, value := range mapUsers {
		fmt.Printf("User %s is %d\n", key, value)
	}
	fmt.Println("--------------------------------------------")
}
