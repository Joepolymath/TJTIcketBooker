package main

import (
	"fmt"
	"sync"
	"ticket-booker/helper"
	"time"
)

// Using package level variables, you cannot use the := syntactical sugar. You've gotta use var or const
const conferenceTickets uint   = 50
var appName = "TJ Go Conference"
var remainingTickets uint = 50
// bookings array/slice
var bookings = make([]UserData, 0)  

// STRUCTS
type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	//  Looping for prompt repeatition
	for remainingTickets > 0 && len(bookings) < 50 {
		// ask user for their name by running the custom function
		firstName, lastName, email, userTickets := getUserInput()

		// input validation
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket( userTickets, firstName, lastName, email)
			// go routine for concurrency
			wg.Add(1) // signifying the number of threads that the main thread should wait for
			go sendTicket( userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
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
		wg.Wait()
	}
}


func greetUsers() {
	fmt.Printf("Welcome to this %s ticket booking application. 🤪\n", appName) 
	fmt.Printf("We have a total of %v tickets and %d are available for booking.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	// using array for-each to extract the first names in each array element
			firstNames := []string{}
			for _, booking := range(bookings) {
				firstNames = append(firstNames, booking.firstName)
			}
			return firstNames			
}


func getUserInput() (string, string, string, uint16) {
		var firstName string
		var lastName string
		var email string
		var userTickets uint16

		fmt.Println("Please enter your first name.")
		fmt.Scan(&firstName) 
		
		fmt.Println("Please enter your last name.")
		fmt.Scan(&lastName) 
		
		fmt.Println("Please enter your email address.")
		fmt.Scan(&email) 
	
		fmt.Println("Please enter number of tickets")
		fmt.Scan(&userTickets) 

		return firstName, lastName, email, userTickets
}

func bookTicket( userTickets uint16,  firstName string, lastName string, email string) {
			remainingTickets = remainingTickets - uint(userTickets)

			// create a map for user
			var userData = UserData {
				firstName: firstName,
				lastName: lastName,
				email: email,
				numberOfTickets: uint(userTickets),
			}

			bookings = append(bookings, userData)
			fmt.Printf("List of booking: %v\n", bookings)
	
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, appName)
}

func sendTicket(userTickets uint16, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("=========================================")
	fmt.Printf("Sending ticket \n%v to email address %v\n", ticket, email )
	fmt.Println("=========================================")
	wg.Done() // removes the thread
}