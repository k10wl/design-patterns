package main

type Circle struct {
	radius int
}

func (el *Circle) accept(v Visitor) int {
	return v.visitForCircle(el)
}

func (el *Circle) getName() string {
	return "Circle"
}
