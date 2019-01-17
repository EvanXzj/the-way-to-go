package main

import (
	"fmt"
	"time"
)

type Address struct {
	HouseNumber      uint32
	Street           string
	HouseNumberAddOn string
	POBox            string
	ZipCode          string
	City             string
	Country          string
}

type VCard struct {
	BirtDate  time.Time
	FirstName string
	LastName  string
	NickName  string
	Photo     string
	Addresses map[string]*Address
}

func main() {
	addr1 := &Address{12, "Elfenstraat", "", "", "2600", "Mechelen", "Belgie"}
	addr2 := &Address{28, "Heideland", "", "", "2640", "Mortsel", "Belgie"}
	addrs := make(map[string]*Address)
	addrs["youth"] = addr1
	addrs["now"] = addr2

	birtdt := time.Date(1956, 1, 17, 15, 4, 5, 0, time.Local)
	photo := "example.png"
	vcard := &VCard{birtdt, "Ivo", "BalBaert", "", photo, addrs}

	fmt.Printf("Here is the full VCard: %v\n", vcard)
	fmt.Printf("My Addresses are:\n%v\n%v\n", addr1, addr2)
}
