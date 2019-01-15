package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "sgo1311你是否@"
	s = strings.Map(mapping, s)

	fmt.Println(s)
}

func mapping(r rune) rune {
	if r > 255 {
		return '?'
	}

	return r
}
