package main

type Director struct {
	builder IBuilder
}

func newDirector(builder IBuilder) Director {
	return Director{builder: builder}
}

func (d *Director) setBuilder(builder IBuilder) {
	d.builder = builder
}

func (d *Director) make() {
	d.builder.reset()
	d.builder.setFrame()
	d.builder.setMotors()
	d.builder.setProcessor()
	d.builder.setCamera()
	d.builder.setBattery()
	d.builder.getDrone()
}

func (d *Director) result() IBuilder {
	return d.builder.getDrone()
}
