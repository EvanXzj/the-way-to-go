package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the server ...")
	// 创建listener
	listener, err := net.Listen("tcp", ":9527")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return // 终止程序
	}
	// 监听并接受来自客户端的链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err)
			return
		}
		fmt.Println("got a connection")
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		length, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		fmt.Printf("Received data: %v\n", string(buf[:length]))
	}
}
