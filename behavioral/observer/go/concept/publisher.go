package main

import (
	"fmt"
	"slices"
)

type Publisher interface {
	subscribe(Subscriber)
	unsubscribe(Subscriber)
	notifySubscribers()
}

type Maganize struct {
	name         string
	latestUpdate int
	subscribers  []Subscriber
}

func newMagazine(name string) *Maganize {
	return &Maganize{name: name, latestUpdate: 0, subscribers: []Subscriber{}}
}

func (m *Maganize) subscribe(s Subscriber) {
	m.subscribers = append(m.subscribers, s)
}

func (m *Maganize) unsubscribe(s Subscriber) {
	m.subscribers = slices.DeleteFunc(m.subscribers, isSubscriber(s))
}

func (m *Maganize) newRelease() {
	m.latestUpdate++
}

func (m *Maganize) notifySubscribers() {
	fmt.Printf("> Sending %s notification\n", m.name)
	for _, s := range m.subscribers {
		s.update(fmt.Sprintf("%s released version %d", m.name, m.latestUpdate))
	}
}

func isSubscriber(s Subscriber) func(e Subscriber) bool {
	return func(e Subscriber) bool {
		return s == e
	}
}
