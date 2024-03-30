package main

import "fmt"

type State struct {
	objects []Context
}

func (s *State) print() {
	for _, object := range s.objects {
		object.flyweight.print(&object.uniqueState)
	}
}

func main() {
	fmt.Println(">>> Flyweight demo")
	factory := newSingleSpriteFactory()
	state := State{}
	state.objects = make([]Context, 10)
	for i := 0; i < len(state.objects); i++ {
		var spriteType string
		if i%2 == 0 {
			spriteType = "hero"
		} else {
			spriteType = "foe"
		}
		flyweight, err := factory.getSprite(spriteType)
		if err != nil {
			panic(err)
		}
		state.objects[i] = Context{
			flyweight: *flyweight,
			uniqueState: ParticlePosition{
				x: i,
				y: i,
			},
		}
	}
	state.print()
}
