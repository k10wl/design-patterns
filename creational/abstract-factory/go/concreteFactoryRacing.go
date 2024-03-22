package main

type RacingFactory struct{}

func (f RacingFactory) makeDrone() IDrone {
	return DroneRacing{
		Drone{
			size:         5,
			batteryLevel: 80,
		},
	}
}

func (f RacingFactory) makeCar() ICar {
	return CarRacing{
		Car{
			size:         2,
			batteryLevel: 80,
		},
	}
}
