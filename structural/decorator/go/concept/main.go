package main

import "fmt"

type IPizza interface {
	getPrice() int
}

type VeggieMania struct{}

func (p *VeggieMania) getPrice() int {
	return 15
}

type AdditionalCheese struct {
	pizza IPizza
}

func (p *AdditionalCheese) getPrice() int {
	return p.pizza.getPrice() + 2
}

type AdditionalMeat struct {
	pizza IPizza
}

func (p *AdditionalMeat) getPrice() int {
	return p.pizza.getPrice() + 2
}

func main() {
	fmt.Println(">>> Decorator demo")

	pizza := &VeggieMania{}
	withCheese := &AdditionalCheese{pizza}
	withCheeseAndMeat := &AdditionalMeat{withCheese}
	fmt.Printf("Total price: %d\n", withCheeseAndMeat.getPrice())
}
