package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint16, remainingTickets uint) (bool, bool, bool) {
		isValidName := len(firstName) > 2 && len(lastName) > 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= uint16(remainingTickets)
		
		return isValidName, isValidEmail, isValidTicketNumber
}