package main

import "fmt"

func main() {
	fmt.Println("Adapter demo:")
	adapter := NewAdapter(&Adaptee{&OldTransmitor{}})
	adapter.transmit("arrived at destination, ready for next commands")
}
