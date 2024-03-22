package main

import "fmt"

type IRCFactory interface {
	makeDrone() IDrone
	makeCar() ICar
}

func getRCFactory(rcType string) (IRCFactory, error) {
	if rcType == "cinematic" {
		return CinematicFactory{}, nil
	}
	if rcType == "racing" {
		return RacingFactory{}, nil
	}
	return nil, fmt.Errorf("unexpected rc type")
}
