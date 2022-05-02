package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("\n")
	fmt.Printf("Welcome to %v booking app \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number for tickets to book: ")
		fmt.Scan(&userTickets)

		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail bool = strings.Contains(email, "@")
		var isValidTicketNum bool = userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNum {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a email at %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			var firstNames = []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of booking are : %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Tickets are sold!")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("Invalid input data. Name too short.\n")
			}
			if !isValidEmail {
				fmt.Printf("Invalid input data. Not an email.\n")
			}
			if !isValidTicketNum {
				fmt.Printf("Invalid input data. Tickets number invalid.\n")
			}
			continue
		}

	}
}
