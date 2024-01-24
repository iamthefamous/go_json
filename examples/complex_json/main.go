package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/sanity-io/litter"
)

type (
	Person struct {
		Name    string  `json:"name"`
		Age     int     `json:"age"`
		Address Address `json:"address"`
		Pets    []Pet   `json:"pets"`
	}

	Address struct {
		Street string `json:"street"`
		City   string `json:"city"`
		State  string `json:"state"`
		Zip    string `json:"zip"`
	}

	Pet struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Type string `json:"type"`
	}
)

func main() {
	b, err := os.ReadFile("assets/complex.json")
	if err != nil {
		log.Fatal("Unable to read of file dut to: ", err)
	}

	var person Person
	err = json.Unmarshal(b, &person)
	if err != nil {
		log.Fatal("Unable to unmarshal doe to: ", err)
	}

	litter.Dump(person)
}
