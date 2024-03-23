package main

import "fmt"

type IBuilder interface {
	reset()
	setBattery()
	setFrame()
	setCamera()
	setMotors()
	setProcessor()
	getDrone() IBuilder
}

func main() {
	manualBuilder := newDroneManualBuilder()
	fpvBuilder := newFPVDroneBuilder()
	director := newDirector(&manualBuilder)
	director.make()
	director.setBuilder(&fpvBuilder)
	director.make()
	fmt.Printf("manual builder result:\n%s\n---\n", manualBuilder.text)
	fmt.Printf("created drone:\n%+v", fpvBuilder.Drone)
}
