package main

import "fmt"

func main() {
	fmt.Println(">>> Chain of responsibility demo")
	ageCheck := NewAgeCheck()
	criminalRecordCheck := NewCriminalRecordCheck()
	inviteListCheck := NewInviteListCheck()
	proceed := NewProceed()
	inviteListCheck.setNext(&ageCheck)
	ageCheck.setNext(&criminalRecordCheck)
	criminalRecordCheck.setNext(&proceed)
	guests := []Guest{
		{name: "Adam", age: 17, criminalRecord: false, invitee: "John"},
		{name: "Jeffry", age: 18, criminalRecord: false, invitee: "John"},
		{name: "Joan", age: 19, criminalRecord: true, invitee: "John"},
		{name: "Anna", age: 28, criminalRecord: false, invitee: "Nancy"},
		{name: "Mary", age: 33, criminalRecord: false, invitee: "Joe"},
		{name: "Kate", age: 17, criminalRecord: false, invitee: "Joe"},
	}
	for _, guest := range guests {
		fmt.Printf("\n\nChecking %s\n", guest.name)
		inviteListCheck.handle(&guest)
	}
}
