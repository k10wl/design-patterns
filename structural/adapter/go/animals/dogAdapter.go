package main

import "fmt"

type DogAdapter struct {
	*Dog
}

func NewDogAdapter(dog *Dog) AnimalAdapter {
	return &DogAdapter{dog}
}

func (d *Dog) Act() {
	fmt.Printf("dog: ")
	d.Woof()
}
