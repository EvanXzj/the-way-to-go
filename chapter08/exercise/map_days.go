package main

import "fmt"

func main() {
	days := map[int]string{
		1: "monday",
		2: "tuesday",
		3: "wednesday",
		4: "thursday",
		5: "friday",
		6: "saturday",
		7: "sunday",
	}

	flagHolliday := false
	for k, v := range days {
		if v == "thursday" || v == "holliday" {
			fmt.Println(v, "is the", k, "th day in the week")

			if v == "holliday" {
				flagHolliday = true
			}
		}
	}

	if !flagHolliday {
		fmt.Println("holliday is not a day!")
	}
}
