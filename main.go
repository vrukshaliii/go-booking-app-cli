package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []string

func main() {

	// calling the greetUser function to greet the users
	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isvalidName, isvalidEmail, isvalidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isvalidName && isvalidEmail && isvalidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			// calls functions print first names
			firstNames := getfirstNames()
			fmt.Printf("First names of bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference tickets are sold out, Come back next year")
			}
		} else {
			if !isvalidName {
				fmt.Println("First name or lastname you entered is too small")
			}

			if !isvalidEmail {
				fmt.Println("Email address you entered you entered doesn't contain @ sign")
			}

			if !isvalidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}

		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask the user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email ID: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isvalidName := len(firstName) >= 2 && len(lastName) >= 2
	isvalidEmail := strings.Contains(email, "@")
	isvalidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isvalidName, isvalidEmail, isvalidTicketNumber
}

func getfirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Slices
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thankyou %v %v for booking %v tickets, you will soon recieve an email from us at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for the conference\n", remainingTickets)
}

// Extras ******ignore*********

// Arrays
// bookings[0] = firstName + " " + lastName

// slice another syntax
// bookings := []string{}

// Switch statements example if you want to add
// city := "London"

// switch city {
// case "Newyork":
// //excute code for booking new york conference tickets
// case "London", "Berlin":
// 	//some code here
// case "Mexico":
// 	//some code
// case "Hong Kong":
// 	//some code
// // case "Berlin":
// // 	//some code
// default:
// 	//some code
// }
// fmt.Printf("The first value: %v\n", bookings[0])
// fmt.Printf("Type of slice %T\n", bookings)
// fmt.Printf("Size of the slice %v\n", len(bookings))
