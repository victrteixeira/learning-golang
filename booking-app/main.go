package main

import "fmt"

func main() {
	// three different ways to declare and assign a value to a variable
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available. ", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	// var book [50]string{"Nana", "Nicole", "Peter"}
	// var bookings [50]string // array declaration -> 50 is the size of array
	// var bookings []string
	bookings := []string{} // slice definition

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for index, booking := range bookings {
			var names = strings.Fields(booking)
			var firstName = names[0] 1:19:
		}

		fmt.Printf("These are all our bookings: %v\n", bookings)
	}
}
