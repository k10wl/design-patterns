package main

import "fmt"

func main() {
	fmt.Println(">>> Mediator demo")
	landing := newLandingSpot()
	plane := Plane{mediator: landing}
	helicopter := Helicopter{mediator: landing}
	plane.land()
	helicopter.land()
	plane.fly()
	plane.land()
	helicopter.fly()
}
