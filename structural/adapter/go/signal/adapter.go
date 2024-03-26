package main

import "fmt"

type TransmitorAdapter interface {
	transmit(binary string)
}

type Adaptee struct {
	*OldTransmitor
}

type Adapter struct {
	*Adaptee
}

func NewAdapter(adaptee *Adaptee) TransmitorAdapter {
	return &Adapter{adaptee}
}

func (adapter *Adapter) transmit(data string) {
	adapter.adaptData(data)
}

func (adaptee *Adaptee) adaptData(data string) {
	fmt.Printf("transforming `%s` to binary...\n", data)
	res := ""
	for _, c := range data {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	adaptee.sendBinary(res)
}
