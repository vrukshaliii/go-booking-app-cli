package main

import (
	"fmt"
	"golang-booking-app/helper"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []UserData


// struct type -> can combine all types of data type for all properties
type UserData struct {
	firstName string
	lastName string
	email string
	numberofTickets uint
}

func main() {

	// calling the greetUser function to greet the users
	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isvalidName, isvalidEmail, isvalidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isvalidName && isvalidEmail && isvalidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			// just use go keyword for concurrency
			go sendTicket(userTickets, firstName, lastName, email)

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

func getfirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		//changed to map so removed the above line
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// using struct type
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberofTickets: userTickets,
	}

	// creating a map for users
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberofTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thankyou %v %v for booking %v tickets, you will soon recieve an email from us at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for the conference\n", remainingTickets)
}


// Example of concurrency
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("##########################################################")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("##########################################################")
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
