package main

type ICar interface {
	getSize() int
	getBatteryLevel() int
}

type Car struct {
	size         int
	batteryLevel int
}

func (c Car) getSize() int {
	return c.size
}

func (c Car) getBatteryLevel() int {
	return c.batteryLevel
}
