package main

import "fmt"

func main() {
	mf := map[int]func() int {
		1: func() int {return 10},
		2: func() int {return 20},
		5: func() int {return 50},
	}

	fmt.Println(mf)	// 打印： map[2:0x1091910 5:0x1091920 1:0x1091900]
}
