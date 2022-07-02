package main

import (
	"encoding/gob"
	"log"
	"os"
)

type Address1 struct {
	Type    string
	City    string
	Country string
}

type VCard1 struct {
	FirstName string
	LastName  string
	Address1  []*Address1
	Remark    string
}

var content string

func main() {
	pa := &Address1{"pivate", "Aartselaar", "Belgium"}
	wa := &Address1{"work", "Boom", "Belgium"}

	vc := VCard1{"Jan", "Kersschot", []*Address1{pa, wa}, "none"}

	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// using an encoder:

	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()

	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)

	if err != nil {
		log.Println("Error in encoding gob")
	}

}
