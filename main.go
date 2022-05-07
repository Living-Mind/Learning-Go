package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	// Validation
	var isValidName, isValidEmail, isValidTicketNum bool = validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNum {

		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		// call function print first names
		var firstNames []string = getFirstNames()
		fmt.Printf("The first names of booking are : %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Tickets are sold!")
			//break
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
		//continue
	}
	wg.Wait()

}

func greetUsers() {

	fmt.Printf("Welcome to %v booking app \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {

	var firstNames = []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v.\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return 0, nil
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// Sleep for process
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("################################")
	wg.Done()
}
