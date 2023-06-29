package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address Address  `json:"address"`
	Phone   []string `json:"phone"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

type ChinaPerson struct {
	*Person
	Province string `json:"province"`
}

func main() {
	jsonStr := `{
        "name": "Alice",
        "age": 30,
        "address": {
            "street": "123 Main St",
            "city": "Anytown"
        },
        "phone": [
            "555-555-1212",
            "555-555-2121"
        ],
		"province": "Guangdong"
    }`

	var person ChinaPerson
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", *person.Person)
}
