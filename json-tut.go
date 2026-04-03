package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address struct {
		Street string `json:"street"`
		City   string `json:"city"`
		State  string	`json:"state"`
	}
	Hobbies []string `json:"hobbies"`
	HasCat  bool     `json:"hasCat"`
}

func main() {
	myJson := `[
	{
	"name":"Mark",
	"age": 30,
	"address":{
		"street":"123 Main St",
		"city":"Anytown",
		"state":"CA"
	},
	"hobbies":["coding","hiking","cooking"],
	"hasCat" :true
},
	{ 
	"name":"Jane",
	"age": 25,
	"address":{
		"street":"456 Elm St",
		"city":"Othertown",
		"state":"NY"
	},
	"hobbies":["painting","traveling","dancing"]
	}
	
	] `

	var unmarshalled []Person

	err:= json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil{
		log.Println("Error unmarshalling json", err)
	}
	log.Printf("unmarshalled %v", unmarshalled)

	// write json from a struct

	var mySlice []Person

	var m1 Person
	m1.Name = "Alice"
	m1.Age = 28
	m1.Address= struct {
		Street string `json:"street"`
		City   string `json:"city"`
		State  string `json:"state"`
	}{
		Street: "789 Oak St",
		City:   "Sometown",
		State:  "TX",
	}
	m1.Hobbies = []string{"reading", "swimming"}
	m1.HasCat = false

	mySlice = append(mySlice,m1)
	var m2 Person
	m2.Name = "Diana"
	m2.Age = 23
	m2.Address= struct {
		Street string `json:"street"`
		City   string `json:"city"`
		State  string `json:"state"`
	}{
		Street: "101 Pine St",
		City:   "Newtown",
		State:  "FL",
	}
	m2.Hobbies = []string{"singing", "dancing"}
	m2.HasCat = false

	mySlice = append(mySlice,m2)

	newJson, err := json.MarshalIndent(mySlice,"", "   ")
	if err != nil{
		log.Println("error, marshalling json", err)
	}
	fmt.Println(string(newJson))
}
