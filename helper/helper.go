package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isvalidName := len(firstName) >= 2 && len(lastName) >= 2
	isvalidEmail := strings.Contains(email, "@")
	isvalidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isvalidName, isvalidEmail, isvalidTicketNumber
}
