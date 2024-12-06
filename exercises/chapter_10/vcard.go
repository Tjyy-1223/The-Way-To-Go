package main

import (
	"fmt"
	"time"
)

type Address struct {
	Street      string
	HouseNumber uint32
}

type VCard struct {
	name     string
	address  map[string]*Address
	birthday time.Time
	photo    string
}

func mainVcard() {
	addr1 := &Address{"Elfenstraat", 12}
	addr2 := &Address{"Heideland", 28}

	addrs := make(map[string]*Address)
	addrs["youth"] = addr1
	addrs["now"] = addr2

	birthdt := time.Date(1956, 1, 17, 15, 4, 5, 0, time.Local)
	photo := "MyDocuments/MyPhotos/photo1.jpg"

	vcard := &VCard{"Ivo", addrs, birthdt, photo}
	fmt.Printf("Here is the full VCard: %v\n", vcard)
	fmt.Printf("My Addresses are: \n %v \n %v \n", addr1, addr2)
}
