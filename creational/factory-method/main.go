package main

import "fmt"

// Golang does not have inheritance and classes, it is impossible to create
// factory without them, following code is simple factory

type IFood interface {
	getSaturaton() int
	getName() string
}

type Food struct {
	saturation int
	name       string
}

type Burger struct {
	Food
}
type Pizza struct {
	Food
}

func getFood(foodType string) (IFood, error) {
	if foodType == "pizza" {
		return newPizza(), nil
	}
	if foodType == "burger" {
		return newBurger(), nil
	}
	return nil, fmt.Errorf("food does not exist")
}

func newBurger() Burger {
	return Burger{
		Food{
			saturation: 1,
			name:       "burger",
		},
	}
}

func (b Food) getSaturaton() int {
	return b.saturation
}

func (b Food) getName() string {
	return b.name
}

func newPizza() Pizza {
	return Pizza{
		Food: Food{
			saturation: 3,
			name:       "pizza",
		},
	}
}

func (p Pizza) getSaturaton() int {
	return p.saturation
}

func (p Pizza) getName() string {
	return p.name
}

func main() {
	pizza, _ := getFood("pizza")
	burger, _ := getFood("burger")

	fmt.Printf("%s satisfies %d\n", pizza.getName(), pizza.getSaturaton())
	fmt.Printf("%s satisfies %d\n", burger.getName(), burger.getSaturaton())
}
