package main

import (
	"fmt"
)

type Proceed struct {
	next CheckChain
}

func NewProceed() Proceed {
	return Proceed{}
}

func (check *Proceed) setNext(next CheckChain) {
	check.next = next
}

func (check Proceed) handle(guest *Guest) {
	fmt.Printf("%s can get in (%+v)\n", guest.name, guest)
}
