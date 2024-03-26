package main

import "fmt"

type (
	Cat  struct{}
	Dog  struct{}
	Fish struct{}
)

func (c *Cat) Meow() {
	fmt.Println("meow!")
}

func (d *Dog) Woof() {
	fmt.Println("woof! woof!")
}

func (f *Fish) SwimInSilance() {
	fmt.Println("I'm not drowning, I'm having fun")
}
