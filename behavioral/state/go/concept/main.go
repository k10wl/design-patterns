package main

import "fmt"

type State interface {
	cancel()
	publish()
}

type Document struct {
	draftState      State
	moderationState State
	adminState      State
	state           State
}

func (d *Document) cancel() {
	fmt.Println("> Sending to state cancel handler")
	d.state.cancel()
}

func (d *Document) publish() {
	fmt.Println("> Sending to current state publish handler")
	d.state.publish()
}

func (d *Document) setState(s State) {
	d.state = s
}

type Draft struct {
	document *Document
}

func (d Draft) publish() {
	fmt.Printf("*draft* Published as post\n")
	d.document.setState(d.document.moderationState)
}

func (d *Draft) cancel() {
	fmt.Printf("*draft* Document deleted forever\n")
}

type Moderator struct {
	document *Document
}

func (m Moderator) publish() {
	fmt.Printf("*moderator* No moderation errors, publishing into public\n")
	m.document.setState(m.document.adminState)
}

func (m Moderator) cancel() {
	fmt.Printf("*moderator* Document hidden by moderator and send to admin review\n")
	m.document.setState(m.document.draftState)
}

type Admin struct {
	document *Document
}

func (a Admin) publish() {
	fmt.Printf("*admin* Admin published\n")
}

func (a Admin) cancel() {
	fmt.Printf("*admin* Marked as censored, send to moderation\n")
	a.document.setState(a.document.moderationState)
}

func main() {
	fmt.Println(">>> State demo")
	d := Document{}
	d.draftState = &Draft{&d}
	d.moderationState = &Moderator{&d}
	d.adminState = &Admin{&d}
	d.setState(d.draftState)
	d.publish()
	d.publish()
	d.publish()
	d.publish()
	d.cancel()
	d.cancel()
	d.cancel()
	d.publish()
	d.publish()
	d.publish()
}
