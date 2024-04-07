package main

import "math"

type AreaCalculator struct{}

func (AreaCalculator) visitForCircle(c *Circle) int {
	return int(math.Floor(math.Pi * (float64(c.radius) * float64(c.radius))))
}

func (AreaCalculator) visitForSquare(s *Square) int {
	return s.side * s.side
}

func (AreaCalculator) visitForRectangle(r *Rectangle) int {
	return r.a * r.b
}
