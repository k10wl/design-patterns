package main

import "fmt"

type ConcreteAdapter struct {
	*Adaptee
}

func NewAdapter(adaptee *Adaptee) Target {
	return &ConcreteAdapter{Adaptee: adaptee}
}

func (adapter *ConcreteAdapter) Operation() {
	fmt.Println("I'm concrete adapter")
	adapter.AdaptedOperation()
}
