package main

import (
	"log"

	"github.com/developermark17/myniceprogram/helpers"
)

func main()  {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Airaz Khan"
	myVar.TypeAge = 22
	log.Println(myVar.TypeName, myVar.TypeAge)
}