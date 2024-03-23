package main

import "strings"

type DroneManual struct {
	text string
}

type DroneManualBuilder struct {
	stringBuilder strings.Builder
	DroneManual
}

func newDroneManualBuilder() DroneManualBuilder {
	return DroneManualBuilder{}
}

func (d *DroneManualBuilder) reset() {
	d.stringBuilder.Reset()
	d = &DroneManualBuilder{}
	d.stringBuilder.WriteString("Drone building manual:\n")
}

func (d *DroneManualBuilder) setBattery() {
	d.stringBuilder.WriteString("disarm drone, strap battery, plug it into the board\n")
}

func (d *DroneManualBuilder) setFrame() {
	d.stringBuilder.WriteString("prepare working surface, place frame facing towards you\n")
}

func (d *DroneManualBuilder) setCamera() {
	d.stringBuilder.WriteString("solder camera wires into controller according to controller instructions. screw camera into frame holder\n")
}

func (d *DroneManualBuilder) setMotors() {
	d.stringBuilder.WriteString("install motors on the frame, ducktape wires to the arms of the frame so they wont wobble\n")
}

func (d *DroneManualBuilder) setProcessor() {
	d.stringBuilder.WriteString("set softening to the frame, and carefully screw processor into the frame\n")
}

func (d *DroneManualBuilder) getDrone() IBuilder {
	d.text = d.stringBuilder.String()
	return d
}
