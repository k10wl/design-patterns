package main

type Director struct {
	builder IDroneBuilder
}

func newDirector(builder IDroneBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) setBuilder(b IDroneBuilder) {
	d.builder = b
}

func (d *Director) buildDrone() Drone {
	d.builder.reset()
	d.builder.setFrame()
	d.builder.setMotors()
	d.builder.setProcessor()
	d.builder.setCamera()
	d.builder.setBattery()
	return d.builder.getDrone()
}
