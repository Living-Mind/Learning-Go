package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("\n")
	fmt.Printf("Welcome to %v booking app \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//ask user for name

	var userName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println(&userName)

	fmt.Println("Enter first name: ")
	fmt.Scan(&userName)

	fmt.Println("Enter last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number for tickets to book: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a email at %v \n", userName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v", remainingTickets, conferenceName)
}
