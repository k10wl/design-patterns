package main

import "fmt"

func main() {
	racingFactory, _ := getRCFactory("racing")
	cinematicFactory, _ := getRCFactory("cinematic")

	rd := racingFactory.makeDrone()
	rc := racingFactory.makeCar()

	cd := cinematicFactory.makeDrone()
	cc := cinematicFactory.makeCar()

	fmt.Println("racing drone stats:", rd.getSize(), rd.getBatterLevel())
	fmt.Println("racing car stats:", rc.getSize(), rc.getBatteryLevel())
	fmt.Println("cinematic drone stats:", cd.getSize(), cd.getBatterLevel())
	fmt.Println("cinematic car stats:", cc.getSize(), cc.getBatteryLevel())
}
