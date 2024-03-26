package main

import "fmt"

type OldTransmitor struct{}

func (oldTransmitor *OldTransmitor) sendBinary(binary string) {
	fmt.Printf("Started transmitting: %v\n", binary)
}
