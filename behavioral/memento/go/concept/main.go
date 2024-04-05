package main

import (
	"fmt"
)

type Originator struct {
	state string
}

func (o *Originator) save() *Memento {
	return newMemento(o.state)
}

func (o *Originator) restore(m *Memento) {
	o.state = m.state
}

func (o *Originator) update(s string) {
	o.state = s
}

func newOriginator() *Originator {
	return &Originator{}
}

type Memento struct {
	state string
}

func newMemento(state string) *Memento {
	return &Memento{state}
}

func (m *Memento) getState() string {
	return m.state
}

type Caretaker struct {
	originator *Originator
	history    []*Memento
}

func (c *Caretaker) updateText(s string) {
	m := c.originator.save()
	// this has bug, history should be overriden after pointer, but whatever
	c.history = append(c.history, m)
	c.originator.update(s)
	fmt.Printf("Writing into history: %s\n", s)
}

func (c *Caretaker) undo() {
	l := len(c.history)
	if l == 0 {
		// no undo available
		return
	}
	fmt.Printf("Undo requested\n")
	last := c.history[l-1]
	c.history = c.history[:l-1]
	c.originator.restore(last)
}

func newCaretaker(originator *Originator) *Caretaker {
	return &Caretaker{
		originator: originator,
		history:    []*Memento{},
	}
}

func main() {
	fmt.Println(">>> Memento demo")

	originator := newOriginator()
	caretaker := newCaretaker(originator)
	caretaker.updateText("Hello")
	caretaker.updateText("How are you?")
	caretaker.updateText("I'm under the water")
	caretaker.updateText("Please help me")
	caretaker.updateText("It's too much raining")
	caretaker.updateText("blbllblbllblblbllb")
	fmt.Printf("state: %+v\n", originator.state)
	caretaker.undo()
	fmt.Printf("state: %+v\n", originator.state)
	caretaker.undo()
	fmt.Printf("state: %+v\n", originator.state)
}
