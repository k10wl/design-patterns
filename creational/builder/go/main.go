package main

import "fmt"

type Drone struct {
	battery   string
	frame     string
	camera    string
	motors    string
	processor string
}

type IDroneBuilder interface {
	reset()
	setBattery()
	setFrame()
	setCamera()
	setMotors()
	setProcessor()
	getDrone() Drone
}

func main() {
	d := newDirector(&FPVDroneBuilder{})
	fpv := d.buildDrone()
	fmt.Printf("created fpv: %+v\n", fpv)
	d.setBuilder(&TinyWhoopDroneBuilder{})
	tinyWhoop := d.buildDrone()
	fmt.Printf("created tiny whoop: %+v\n", tinyWhoop)
}
