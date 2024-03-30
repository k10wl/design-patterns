package main

import "fmt"

type Sprite struct {
	background string
}

func (s *Sprite) print(p *ParticlePosition) {
	fmt.Printf(
		"Drawing flyweight on canvas: \n- x %d\n-y %d\n- sprite %s\n",
		p.x,
		p.y,
		s.background,
	)
}
