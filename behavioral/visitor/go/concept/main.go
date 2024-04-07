package main

import "fmt"

type Visitor interface {
	visitForCircle(*Circle) int
	visitForSquare(*Square) int
	visitForRectangle(*Rectangle) int
}

func main() {
	fmt.Println(">>> Visitor demo")
	c := Circle{radius: 2}
	s := Square{side: 2}
	r := Rectangle{a: 3, b: 4}
	calc := AreaCalculator{}
	fmt.Printf("%s area: %d\n", c.getName(), c.accept(calc))
	fmt.Printf("%s area: %d\n", s.getName(), s.accept(calc))
	fmt.Printf("%s area: %d\n", r.getName(), r.accept(calc))
}
