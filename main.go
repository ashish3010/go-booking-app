package main

import (
	"fmt"
	"sync"
)

var conferenceName string = "Go conference"

const conferenceTicket uint = 50 // contstant
var remainingTickets uint = 50
var bookings = make([]UserData, 0)
var wg = sync.WaitGroup{}

func main() {
	GreetUser()

	// for {
	firstName, lastName, email, userTickets := GetUserInput()
	isNameValid, isEmailValid, isValidTicketCount := CheckValidity(firstName, lastName, email, userTickets)

	if isNameValid && isEmailValid && isValidTicketCount {
		// if remainingTickets <= 0 {
		// 	fmt.Println("Sorry, all tickets are sold.")
		// 	break
		// }
		BookTicket(userTickets, firstName, lastName, email)
	} else if remainingTickets == 0 {
		fmt.Println("Sorry, all tickets are sold.")
	} else {
		ShowValidMessage(isEmailValid, isNameValid, isValidTicketCount, userTickets)
	}
	// wg.Wait() will wait for all the secondary threads to complete before completing the main thread
	wg.Wait()
	// }

}
