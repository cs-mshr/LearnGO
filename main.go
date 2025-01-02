package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	myJson := `[
		{
			"name": "John", 
			"age": 30
		},
		{
			"name": "Jane",
			"age": 25
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		panic(err)
	}

	log.Println(unmarshalled)

	var mySlice []Person
	m1 := Person{
		Name: "akshat",
		Age:  23,
	}
	m2 := Person{
		Name: "ajay",
		Age:  21,
	}

	for i := 0; i < 2; i++ {
		mySlice = append(mySlice, m1, m2)
	}

	//newJson, err := json.Marshal(mySlice)
	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		panic(err)
	}
	log.Println(string(newJson))
}
