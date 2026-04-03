 package main

import "fmt"


func main() {
	ifElsecondition()
}

func ifElsecondition(){
	age := 18

	if(age<14){
		fmt.Println("You are a minor")
	}else if(age>=14 && age<18){
		fmt.Println("You are a teenager")
	}
}
