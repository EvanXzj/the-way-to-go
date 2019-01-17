package main

import (
	"fmt"
	"reflect"
)

// 结构体中的字段除了有名字和类型外，还可以有一个可选的标签（tag）：它是一个附属于字段的字符串，可以是文档或其他的重要标记
type TagType struct {
	f1 bool   "An important answer"
	f2 string "The name of the thing"
	f3 int    "How much there are"
}

func main() {
	tt := TagType{false, "XX", 1}
	for i := 0; i < 3; i++ {
		reflectTag(tt, i)
	}
}

func reflectTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	fmt.Println(ttType)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
