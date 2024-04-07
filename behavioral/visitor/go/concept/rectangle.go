package main

type Rectangle struct {
	a int
	b int
}

func (el *Rectangle) accept(v Visitor) int {
	return v.visitForRectangle(el)
}

func (el *Rectangle) getName() string {
	return "Rectangle"
}
