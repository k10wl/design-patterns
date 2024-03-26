package main

type AnimalAdapter interface {
	Act()
}

type Adapter struct {
	adaptee AnimalAdapter
}

func NewAdapter(adaptee AnimalAdapter) AnimalAdapter {
	return Adapter{adaptee: adaptee}
}

func (a Adapter) Act() {
	a.adaptee.Act()
}
