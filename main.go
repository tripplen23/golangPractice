package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0) // create an empty list of map, 0 is initial size of slice

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	// TODO: Call the function for greeting the user
	greetUsers()

	// TODO: Call the function for getting the user input
	firstName, lastName, email, userTickets := getUserInput()

	// TODO: Call the function for Validating the user input
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		//TODO: Call function for booking ticket
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		//TODO: Call function for notifying to send tickets to user's email address
		go sendTicket(userTickets, firstName, lastName, email)

		//TODO: Call function for returning firstNames
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// TODO: end the program
			fmt.Println("Our conference is booked out. Come back next year")
			// break
		}

	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short.")
		}
		if !isValidEmail {
			fmt.Println("Your email address is not correct.")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid.")
		}
	}
	wg.Wait()
}

// TODO: Greeting
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets, and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

// TODO: User input
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// Ask user for their information
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// TODO: To print out ticket's information
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//TODO: Add the next element to the slice
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v.\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("There are still %v tickets for %v available in the system.\n",
		remainingTickets, conferenceName)
}

// TODO: Print first names of booking list
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#############")
	fmt.Printf("Sending ticket:\n %v \nto email address: %v\n", ticket, email)
	fmt.Println("#############")
	wg.Done()
}
