package main

import "fmt"

func main() {
	fmt.Println("> Adapter demo")
	adapters := []AnimalAdapter{
		NewAdapter(NewCatAdapter(&Cat{})),
		NewAdapter(NewDogAdapter(&Dog{})),
		NewAdapter(NewFishAdapter(&Fish{})),
	}
	for _, a := range adapters {
		a.Act()
	}
}
