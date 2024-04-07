package main

type Square struct {
	side int
}

func (el *Square) accept(v Visitor) int {
	return v.visitForSquare(el)
}

func (el *Square) getName() string {
	return "Square"
}
