package main

import (
	"fmt"
	"strings"
	"time"
)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	ticketCount uint
}

func ShowValidMessage(isEmailValid bool, isNameValid bool, isValidTicketCount bool, userTickets uint) {

	if !isEmailValid {
		fmt.Println("Your email is wrong")
	}
	if !isNameValid {
		fmt.Println("First and last name should have minimum 2 characters")
	}
	if !isValidTicketCount {
		fmt.Printf("Apologies, but you can not book %v tickets. Because we have only %v tickets left", userTickets, remainingTickets)
	}
}

func CheckValidity(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isNameValid := len(firstName) >= 2 && len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@") && strings.Contains(email, ".com")
	isValidTicketCount := userTickets <= remainingTickets

	return isNameValid, isEmailValid, isValidTicketCount
}

func GetUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("How many tickets do you want: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func GreetUser() {
	fmt.Printf("Welcome to our %v booking app \n", conferenceName)
	fmt.Println("Get your tickets here to attend")
}

func BookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		ticketCount: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("%v %v booked %v tickets, and we have also sent the copy of tickets on %v \n", firstName, lastName, userTickets, email)

	fmt.Printf("We have %v tickets and now %v tickets are available \n", conferenceTicket, remainingTickets)

	fmt.Println("all bookings", bookings)

	// wg.Add() will add the second thread to the main thread
	// wg.Add() takes the number of secondary threads that main have to wait
	wg.Add(1)
	//  go command makes line of code concurrent
	// it will not wait to send the ticket
	// it will send the ticket after 10 sec like setTimeout
	go sendTicket(userTickets, firstName, lastName, email)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("Sending ticket: \n %v for %v %v \n to email address %v", userTickets, firstName, lastName, email)
	fmt.Println("#######################")
	fmt.Println(ticket)
	fmt.Println("#######################")
	// wg.Done() will remove the second thread after complition of the thread
	wg.Done()
}
