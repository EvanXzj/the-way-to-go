package main

import (
	"fmt"
	"container/list"
)

func main() {
	l := list.New()
	l.PushBack(100)
	l.PushBack(101)
	l.PushBack(102)

	for e := l.Front(); e != nil; e = e.Next() {
		// fmt.Println(e) 能告诉我这输出的是个什么东西吗？
		fmt.Println(e.Value)
	}
}

