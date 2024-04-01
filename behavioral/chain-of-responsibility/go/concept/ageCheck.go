package main

import "fmt"

type AgeCheck struct {
	next CheckChain
}

func NewAgeCheck() AgeCheck {
	return AgeCheck{}
}

func (check *AgeCheck) setNext(next CheckChain) {
	check.next = next
}

func (check AgeCheck) handle(guest *Guest) {
	if guest.age < 18 {
		fmt.Printf("no underaged, get out, %s, you are %d (%+v)\n", guest.name, guest.age, guest)
		return
	}
	fmt.Printf("%s, age check is ok\n", guest.name)
	if check.next != nil {
		check.next.handle(guest)
	}
}
