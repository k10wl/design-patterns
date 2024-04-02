package main

type Paster interface {
	fromBuffer() string
}

type PasteReceiver struct{}

func (PasteReceiver) fromBuffer() string {
	return "**string from buffer**"
}
