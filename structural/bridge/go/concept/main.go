package main

import (
	"fmt"
	"math"
)

type TV struct {
	power   bool
	channel int
	volume  int
}

func (tv *TV) isEnabled() bool {
	return tv.power
}

func (tv *TV) enable() {
	tv.power = true
}

func (tv *TV) disable() {
	tv.power = false
}

func (tv *TV) getVolume() int {
	return tv.volume
}

func (tv *TV) setVolume(volume int) {
	tv.volume = volume
}

func (tv *TV) getChannel() int {
	return tv.channel
}

func (tv *TV) setChannel(channel int) {
	tv.channel = channel
}

type Device interface {
	isEnabled() bool
	enable()
	disable()
	getVolume() int
	setVolume(int)
	getChannel() int
	setChannel(int)
}

type IRemote interface {
	togglePower()
	channelDown()
	channelUp()
	volumeDown()
	volumeUp()
}

type Remote struct {
	device Device
}

func NewRemote(device Device) IRemote {
	return &Remote{device}
}

func (r *Remote) togglePower() {
	if r.device.isEnabled() {
		r.device.disable()
	} else {
		r.device.enable()
	}
}

func (r *Remote) channelDown() {
	c := r.device.getChannel() - 1
	if c < 0 {
		c = 10
	}
	r.device.setChannel(c)
}

func (r *Remote) channelUp() {
	c := r.device.getChannel() + 1
	if c > 10 {
		c = 1
	}
	r.device.setChannel(c)
}

func (r *Remote) volumeDown() {
	c := int(math.Max(float64(r.device.getVolume()-10), 0))
	r.device.setVolume(c)
}

func (r *Remote) volumeUp() {
	c := int(math.Min(float64(r.device.getVolume()+10), 100))
	r.device.setVolume(c)
}

func main() {
	fmt.Println("> Bridge demo")
	tv := TV{power: false, volume: 50, channel: 1}
	fmt.Printf("tv state: %+v\n", tv)
	remote := NewRemote(&tv)
	remote.togglePower()
	fmt.Printf("after toggle power: %+v\n", tv)
	remote.channelDown()
	fmt.Printf("chan down: %+v\n", tv)
	remote.channelDown()
	fmt.Printf("chan down: %+v\n", tv)
	remote.volumeUp()
	fmt.Printf("volume up: %+v\n", tv)
	remote.volumeUp()
	fmt.Printf("volume up: %+v\n", tv)
	remote.togglePower()
	fmt.Printf("turn off: %+v\n", tv)
}
