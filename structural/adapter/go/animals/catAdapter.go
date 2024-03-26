package main

import "fmt"

type CatAdapter struct {
	*Cat
}

func NewCatAdapter(cat *Cat) AnimalAdapter {
	return &CatAdapter{cat}
}

func (c *Cat) Act() {
	fmt.Printf("cat: ")
	c.Meow()
}
