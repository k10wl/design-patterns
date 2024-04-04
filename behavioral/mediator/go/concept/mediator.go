package main

type Mediator interface {
	initiateLanding(FlyingVehicle) bool
	notifyDeparture()
}

type LandingSpot struct {
	queue  []FlyingVehicle
	isFree bool
}

func newLandingSpot() *LandingSpot {
	return &LandingSpot{
		isFree: true,
	}
}

func (l *LandingSpot) initiateLanding(fl FlyingVehicle) bool {
	if l.isFree {
		l.isFree = false
		return true
	}
	l.queue = append(l.queue, fl)
	return false
}

func (l *LandingSpot) notifyDeparture() {
	if !l.isFree {
		l.isFree = true
	}
	if len(l.queue) == 0 {
		return
	}
	head := l.queue[0]
	head.land()
	l.queue = l.queue[1:]
}
