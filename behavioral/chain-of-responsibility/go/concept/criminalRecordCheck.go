package main

import "fmt"

type CriminalRecordCheck struct {
	next CheckChain
}

func NewCriminalRecordCheck() CriminalRecordCheck {
	return CriminalRecordCheck{}
}

func (check *CriminalRecordCheck) setNext(next CheckChain) {
	check.next = next
}

func (check CriminalRecordCheck) handle(guest *Guest) {
	if guest.criminalRecord {
		fmt.Printf("no criminals, get out %s (%+v)\n", guest.name, guest)
		return
	}
	fmt.Printf("%s, criminal record check is ok\n", guest.name)
	if check.next != nil {
		check.next.handle(guest)
	}
}
