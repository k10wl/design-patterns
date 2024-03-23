package main

type Point struct {
	x int
	y int
}

func newPoint(x int, y int) Point {
	return Point{x: x, y: y}
}

func (p Point) clone() IPrototype {
	return Point{
		x: p.x + 1,
		y: p.y + 1,
	}
}
