/*
*	注意 map 不是按照 key 的顺序排列的，也不是按照 value 的序排列的。
*/

package main

import "fmt"

func main() {
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	for key, value := range map1 {
		fmt.Printf("Key is: %d - value is: %f\n", key, value)
	}

	capitals := map[string] string {"France":"Paris", "Italy":"Rome", "Japan":"Tokyo" }
for key := range capitals {
    fmt.Println("Map item: Capital of", key, "is", capitals[key])
}
}
