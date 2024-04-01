package main

import (
	"fmt"
	"slices"
)

type InviteListCheck struct {
	next CheckChain
}

var inviteList []string = []string{"Joe", "John"}

func NewInviteListCheck() InviteListCheck {
	return InviteListCheck{}
}

func (check *InviteListCheck) setNext(next CheckChain) {
	check.next = next
}

func (check InviteListCheck) handle(guest *Guest) {
	if slices.Index(inviteList, guest.invitee) == -1 {
		fmt.Printf("%s is not on the list, get out (%+v)\n", guest.name, guest)
		return
	}
	fmt.Printf("%s, invite list check is ok\n", guest.name)
	if check.next != nil {
		check.next.handle(guest)
	}
}
