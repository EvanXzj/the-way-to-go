package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s\n", js)
	// Using an encoder:
	file, _ := os.OpenFile("./chapter12/examples/vcard.json", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		fmt.Printf("Error in encoding json\n")
	}

	var newV VCard
	err = json.Unmarshal(js, &newV)
	if err != nil {
		fmt.Printf("Error in Unmarshal json\n")
	}
	fmt.Println(newV, *newV.Addresses[0], *newV.Addresses[1])

	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	err = json.Unmarshal(b, &f)
	if err != nil {
		fmt.Printf("Error in Unmarshal json\n")
	}
	fmt.Println(f)
	for k, v := range f.(map[string]interface{}) {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don’t know how to handle")
		}
	}
}

// Age is of a type I don’t know how to handle. Why ????
