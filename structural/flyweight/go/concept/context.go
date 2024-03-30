package main

type Context struct {
	uniqueState ParticlePosition
	flyweight   Sprite
}

func (c *Context) print() {
	c.flyweight.print(&c.uniqueState)
}
