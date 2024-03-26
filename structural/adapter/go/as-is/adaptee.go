package main

import "fmt"

type Adaptee struct{}

func (adaptee *Adaptee) AdaptedOperation() {
	fmt.Println("I am AdaptedOperation()")
}
