package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Breed         string
	Name          string
	FavoriteTreat string
	Age           int
}

func main() {
	input := `{
		"Breed": "Dachshund",
		"Name": "Frankie",
		"FavoriteTreat": "Chewy sticks",
		"Age": 3
	}`

	var dog Dog

	err := json.Unmarshal([]byte(input), &dog)
	if err != nil {
		log.Fatalf("Unable to unmarshal JSON due to : %s", err)
	}

	fmt.Printf("%s is a %d year old %s who loves %s\n",
		dog.Name,
		dog.Age,
		dog.Breed,
		dog.FavoriteTreat)
}
