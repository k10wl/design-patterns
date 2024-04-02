package main

import "fmt"

type Shortcut struct {
	command Command
}

type Button struct {
	command Command
}

func (b Button) execute() {
	fmt.Println("executing button command")
	b.command.execute()
}

func (s Shortcut) execute() {
	fmt.Println("executing shortcut command")
	s.command.execute()
}
