package main

import "fmt"

func main() {
	const conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}
