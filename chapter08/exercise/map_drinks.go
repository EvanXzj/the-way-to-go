package main

import (
	"fmt"
	"sort"
)

func main() {
	drinks := map[string]string{
		"beer":   "bière",
		"wine":   "vin",
		"water":  "eau",
		"coffee": "café",
		"thea":   "thé",
	}

	sdrinks := make([]string, len(drinks))
	ix := 0

	fmt.Println("The following drinks are avaiable:")
	for eng := range drinks {
		sdrinks[ix] = eng
		ix++

		fmt.Println(eng)
	}

	fmt.Println()
	for eng, fr := range drinks {
		fmt.Printf("The french for %s is %s\n", eng, fr)
	}

	// Sorting:
	fmt.Println()
	fmt.Println("Now the sorted output:")
	sort.Strings(sdrinks)
	
	fmt.Printf("The following sorted drinks are available:\n")
	for _, eng := range sdrinks {
		fmt.Println(eng)
	}
	
	fmt.Println()
	for _, eng := range sdrinks {
		fmt.Printf("The french for %s is %s\n", eng, drinks[eng])
	}
}
