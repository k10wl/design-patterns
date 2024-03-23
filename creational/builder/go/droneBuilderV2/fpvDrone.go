package main

type Drone struct {
	battery   string
	frame     string
	camera    string
	motors    string
	processor string
}

type FPVDroneBuilder struct {
	Drone
}

func newFPVDroneBuilder() FPVDroneBuilder {
	return FPVDroneBuilder{}
}

func (d *FPVDroneBuilder) reset() {
	d = &FPVDroneBuilder{}
}

func (d *FPVDroneBuilder) setBattery() {
	d.battery = "20v"
}

func (d *FPVDroneBuilder) setFrame() {
	d.frame = "5inch"
}

func (d *FPVDroneBuilder) setCamera() {
	d.camera = "DJI"
}

func (d *FPVDroneBuilder) setMotors() {
	d.motors = "20000rpm"
}

func (d *FPVDroneBuilder) setProcessor() {
	d.processor = "AIO"
}

func (d *FPVDroneBuilder) getDrone() IBuilder {
	return d
}
