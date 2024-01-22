package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	input := `{
		"name" : "Wednesday",
		"age" : 6,
		"parents" : ["Gomez","Morticia"]
	}`

	var target map[string]any

	err := json.Unmarshal([]byte(input), &target)
	if err != nil {
		log.Fatal("Unable to unmarshal JSON due to : ", err)
	}

	for k, v := range target {
		fmt.Printf("Key: %s, Value: %v\n", k, v)
	}

}
