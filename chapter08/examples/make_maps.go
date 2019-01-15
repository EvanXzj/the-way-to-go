/*
*  1. map 是引用类型, map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节，无论实际上存储了多少数据.
*  2. 未初始化的map的值是nil
*  3. 常用的 len(map1) 方法可以获得 map 中的 pair 数目
*/

package main

import "fmt"

func main() {
	var mapLit map[string]int
	var mapAssigned map[string]int

	// 初始化
	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.1415925
	mapCreated["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
    fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
    fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
	
	fmt.Println(mapAssigned)
	mapLit["one"] = 11
	fmt.Println(mapLit, mapAssigned)
}
