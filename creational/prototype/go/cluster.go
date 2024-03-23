package main

import "fmt"

type Cluster struct {
	name    string
	content []IPrototype
}

func newCluser(name string, content []IPrototype) Cluster {
	return Cluster{name: name, content: content}
}

func (c Cluster) clone() IPrototype {
	copy := Cluster{}
	copy.name = fmt.Sprintf("%s_clone", c.name)
	copy.content = make([]IPrototype, len(c.content))
	for i, p := range c.content {
		copy.content[i] = p.clone()
	}
	return copy
}
