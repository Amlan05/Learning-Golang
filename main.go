package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 30
var bookings = make([]userData, 0) //we can also create slice with this syntax

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidUserTickets := validateUserInput(firstName, lastName, userTickets, email, remainingTickets)

	if isValidName && isValidEmail && isValidUserTickets {

		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		fmt.Printf("The first names of the bookings list are: %v\n", getFirstNames())

	} else if remainingTickets == 0 {
		fmt.Println("Our conference is booked out. Come back next Year")
		// break
	} else {
		if isValidName == false {
			fmt.Println("Enter Valid Name")
		}
		if isValidEmail == false {
			fmt.Println("Invalid EmailId")
		}
		if isValidUserTickets == false {
			fmt.Println("Enter valid no. of Tickets")
		}

		fmt.Println("Your Input data is invalid, try again")
	}

	wg.Wait()

}

// greeting users
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("The conference Tickets is of type %T and conference name is of type %T\n", conferenceTickets, conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// getting first names
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// Getting user Input
func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your first name:")
	fmt.Scan(&firstName) //saves user's value in "userName" variable
	fmt.Printf("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Printf("Enter EmailId:")
	fmt.Scan(&email)

	fmt.Printf("Enter no. of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// booking ticket
func bookTickets(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets -= userTickets

	//Creating a map for a user
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a conformation email on %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining tickets are %v for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v to email %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
