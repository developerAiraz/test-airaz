package main

import (
	"log"

	"github.com/developermark17/myniceprogram/helpers"
)

const numPool = 10;
func Calculatevalue(intChan chan int){
	randomNumber:= helpers.RandomNumber(numPool)
	intChan <- randomNumber
}
func main(){
	intChan := make(chan int)
	defer close(intChan)

	go Calculatevalue(intChan)
	num:= <-intChan
	log.Println(num)
}
