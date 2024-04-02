package main

import "fmt"

type PasteCommand struct {
	receiver Paster
}

func (pr PasteCommand) execute() {
	fmt.Printf("pasted: %s\n", pr.receiver.fromBuffer())
}
