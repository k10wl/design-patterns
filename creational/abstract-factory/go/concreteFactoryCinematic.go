package main

type CinematicFactory struct{}

func (f CinematicFactory) makeDrone() IDrone {
	return DroneCinematic{
		Drone{
			size:         5,
			batteryLevel: 100,
		},
	}
}

func (f CinematicFactory) makeCar() ICar {
	return CarCinematic{
		Car{
			size:         10,
			batteryLevel: 100,
		},
	}
}
