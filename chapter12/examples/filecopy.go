package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CopyFile("./chapter12/examples/target.txt", "./chapter12/examples/source.txt")
	fmt.Println("Copy done!")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
