package main

import "fmt"

type Strategy interface {
	execute(string)
}

type User struct {
	currentDestination string
	roadmap            []string
	strategy           Strategy
}

func (u *User) setStrategy(s Strategy) {
	u.strategy = s
}

func (u *User) travel() {
	u.currentDestination = u.roadmap[0]
	u.roadmap = u.roadmap[:1]
	u.strategy.execute(u.currentDestination)
}

type Bike struct{}

func (b *Bike) execute(s string) {
	fmt.Printf("Biking to %s\n", s)
}

type Car struct{}

func (c *Car) execute(s string) {
	fmt.Printf("Using car to get to %s\n", s)
}

func main() {
	fmt.Println(">>> Strategy demo")
	user := User{roadmap: []string{"Neighbor State", "Local grocery store"}}
	user.setStrategy(&Car{})
	user.travel()
	user.setStrategy(&Bike{})
	user.travel()
}
