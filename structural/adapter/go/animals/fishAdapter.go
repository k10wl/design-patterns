package main

import "fmt"

type FishAdapter struct {
	*Fish
}

func NewFishAdapter(fish *Fish) AnimalAdapter {
	return &FishAdapter{fish}
}

func (f *Fish) Act() {
	fmt.Printf("fish: ")
	f.SwimInSilance()
}
