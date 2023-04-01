package main

import (
	"fmt"
	"strings"
)

func main() {
	appName := "TJ Go Conference"
	const conferenceTickets uint   = 50
	var remainingTickets uint = 50
	// bookings array/slice
	bookings := []string{}
	//  or var bookings []string
	 
	fmt.Printf("Welcome to this %s ticket booking application. 🤪\n", appName) 
	fmt.Printf("We have a total of %v tickets and %d are available for booking.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//  Looping for prompt repeatition
	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint16
		
		// ask user for their name
		fmt.Println("Please enter your first name.")
		fmt.Scan(&firstName) 
		
		fmt.Println("Please enter your last name.")
		fmt.Scan(&lastName) 
		
		fmt.Println("Please enter your email address.")
		fmt.Scan(&email) 
	
		fmt.Println("Please enter number of tickets")
		fmt.Scan(&userTickets) 

		// input validation
		isValidName := len(firstName) > 2 && len(lastName) > 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets < uint16(remainingTickets)


		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - uint(userTickets)
			bookings = append(bookings, firstName + " " + lastName)
	
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, appName)

			// using array for-each to extract the first names in each array element
			firstNames := []string{}
			for _, booking := range(bookings) {
				splittedName := strings.Fields(booking)
				firstName := splittedName[0]
				firstNames = append(firstNames, firstName)
			}
			
			fmt.Printf("These are all our bookings: %v\n", firstNames)
			if remainingTickets < 1 {
				// break the loop here
				fmt.Printf("Our remaining tickets are booked or you are booking more than what is left. %v tickets remaining", remainingTickets)
				break
			} 
		} else {
			if !isValidName {
				fmt.Println("At least, a name you entered is too short. Can only allow names with more than 2 characters")
			}
			if !isValidEmail {
				fmt.Println("Email does not contain '@' character")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
}