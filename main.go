package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// go datatypes tutorial

	var bookings [50]string

	var firstName string
	var lastName string
	var userTickets int
	// ask user for their name
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	for i := 0; i < userTickets; i++ {
		bookings[i] = fmt.Sprintf("%v %v", firstName, lastName)
		io.WriteString(os.Stdout, fmt.Sprintf("%vnd Book of %v \n", i+1, bookings[i]))
	}

	fmt.Printf("Thank you %v %v for booking %v tickets. \n", firstName, lastName, userTickets)
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)
}
