package main

import (
	"encoding/json"
	"fmt"
	"gobasic/customer"
)

func main() {
	// Define the map
	map1 := map[string]interface{}{
		"name": "Some User",
		"age":  35,
		"address": map[string]interface{}{
			"street": "Random St",
			"city":   "Some town",
			"state":  "Some state",
			"zip":    123,
		},
	}

	// Convert the map to JSON
	if jsonContent, err := json.MarshalIndent(map1, "", "    "); err != nil {
		fmt.Println(err)
		return
	} else {
		// Print the JSON data
		fmt.Println(string(jsonContent))
	}

	fmt.Println(customer.Name)
}
