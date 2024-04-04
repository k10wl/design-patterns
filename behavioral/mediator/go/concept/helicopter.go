package main

import "fmt"

type Helicopter struct {
	mediator Mediator
}

func (h *Helicopter) fly() {
	fmt.Println("helicopter is airborne")
	h.mediator.notifyDeparture()
}

func (h *Helicopter) land() {
	if !h.mediator.initiateLanding(h) {
		fmt.Println("helicopter hovering")
		return
	}
	fmt.Println("helicopter landed")
}
