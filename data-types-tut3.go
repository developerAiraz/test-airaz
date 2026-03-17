package main

import "fmt"

func main() {
dataTypes()
}

func dataTypes() {

	var userName string
	var userAge int
	// ask user for their name

	userName = "Mark"
	userAge = 27;
	var userMobile  = 1234567890
	var userEmail  = "mark@example.com"

	fmt.Println(userName)
	fmt.Println(userAge)

	fmt.Printf("User mobile is %T, and user email is %T\n",userMobile, userEmail)
	fmt.Printf("User name is %v and user age is %v\n",userName,userAge)
}
