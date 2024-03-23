package main

type TinyWhoopDroneBuilder struct {
	Drone
}

func (d *TinyWhoopDroneBuilder) reset() {
	d = &TinyWhoopDroneBuilder{}
}

func (d *TinyWhoopDroneBuilder) setBattery() {
	d.battery = "4.2v"
}

func (d *TinyWhoopDroneBuilder) setFrame() {
	d.frame = "3inch"
}

func (d *TinyWhoopDroneBuilder) setCamera() {
	d.camera = "analog"
}

func (d *TinyWhoopDroneBuilder) setMotors() {
	d.motors = "12000rpm"
}

func (d *TinyWhoopDroneBuilder) setProcessor() {
	d.processor = "AIO mini"
}

func (d *TinyWhoopDroneBuilder) getDrone() Drone {
	return d.Drone
}
