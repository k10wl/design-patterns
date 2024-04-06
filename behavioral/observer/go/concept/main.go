package main

import "fmt"

type Subscriber interface {
	update(string)
}

type User struct {
	name string
}

func (u *User) update(s string) {
	fmt.Printf("%s received an update from %s\n", u.name, s)
}

func main() {
	fmt.Println(">>> Observer demo")
	mark := User{"Mark"}
	victor := User{"Victor"}
	sienna := User{"Sienna"}
	kerillian := User{"Kerillian"}
	mayflyCaretaker := newMagazine("Mayfly Caretaker")
	humanTimes := newMagazine("Human Times")
	mayflyCaretaker.subscribe(&kerillian)
	mayflyCaretaker.newRelease()
	humanTimes.subscribe(&mark)
	humanTimes.subscribe(&victor)
	humanTimes.subscribe(&sienna)
	humanTimes.newRelease()
	humanTimes.notifySubscribers()
	mayflyCaretaker.newRelease()
	mayflyCaretaker.newRelease()
	mayflyCaretaker.newRelease()
	mayflyCaretaker.notifySubscribers()
	humanTimes.unsubscribe(&sienna)
	mayflyCaretaker.subscribe(&sienna)
	humanTimes.newRelease()
	humanTimes.notifySubscribers()
	mayflyCaretaker.newRelease()
	mayflyCaretaker.newRelease()
	mayflyCaretaker.notifySubscribers()
}
