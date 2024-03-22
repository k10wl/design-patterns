package main

type IDrone interface {
	getSize() int
	getBatterLevel() int
}

type Drone struct {
	size         int
	batteryLevel int
}

func (d Drone) getSize() int {
	return d.size
}

func (d Drone) getBatterLevel() int {
	return d.batteryLevel
}
