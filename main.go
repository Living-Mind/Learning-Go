package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings []string

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		// Validation
		var isValidName, isValidEmail, isValidTicketNum bool = validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNum {

			bookTicket(userTickets, firstName, lastName, email)

			// call function print first names
			var firstNames []string = getFirstNames()
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

func greetUsers() {

	fmt.Printf("Welcome to %v booking app \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {

	var firstNames = []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {

	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicketNum bool = userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNum
}

func getUserInput() (string, string, string, uint) {

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

	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) (uint, []string) {

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return remainingTickets, bookings

}
