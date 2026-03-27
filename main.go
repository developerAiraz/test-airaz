package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var Firstname string
		var LastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Println("Enter your first name:")
		fmt.Scan(&Firstname)

		fmt.Println("Enter your last name:")
		fmt.Scan(&LastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter your tickets:")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets{
			fmt.Printf("We only have %v tickets remaining , so you can't book %v tickets\n", remainingTickets, userTickets)
			continue

		}
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, Firstname+" "+LastName)

		fmt.Printf("The slice array: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Type of Slice: %T\n", bookings[0])
		fmt.Printf("Size of Slice: %v\n", len((bookings)))
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
			Firstname, LastName, userTickets, email)
		fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}

		for _, booking := range bookings{
		var names = strings.Fields(booking)
		firstNames = append(firstNames,names[0])
		}
		fmt.Printf("The first names of bookings are : %v\n", firstNames)

		if remainingTickets ==0{
		fmt.Print("tickets are sold out,Come back next time");
		break;
	}
	}

	
	
}
