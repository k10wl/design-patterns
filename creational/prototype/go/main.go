package main

import (
	"fmt"
)

type IPrototype interface {
	clone() IPrototype
}

func main() {
	a := newCluser("main", []IPrototype{newPoint(1, 1), newPoint(10, 10), newPoint(20, 20)})

	fmt.Printf("cluster: %+v\n", a)
	fmt.Printf("cluster clone: %+v\n", a.clone())
}
