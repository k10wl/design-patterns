package main

import "fmt"

type Plane struct {
	mediator Mediator
}

func (p *Plane) fly() {
	fmt.Println("plane is airborne")
	p.mediator.notifyDeparture()
}

func (p *Plane) land() {
	if !p.mediator.initiateLanding(p) {
		fmt.Println("plane circling around landing spot")
		return
	}
	fmt.Println("plane landed")
}
