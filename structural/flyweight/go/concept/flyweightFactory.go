package main

import "fmt"

type SpriteFactory struct {
	spriteMap map[string]*Sprite
}

var single = &SpriteFactory{
	spriteMap: map[string]*Sprite{
		"hero": {background: "hero.png"},
		"foe":  {background: "foe.png"},
	},
}

func newSingleSpriteFactory() *SpriteFactory {
	return single
}

func (s *SpriteFactory) getSprite(name string) (*Sprite, error) {
	val, ok := s.spriteMap[name]
	if !ok {
		return nil, fmt.Errorf("sprite with name %s does not exits", name)
	}
	fmt.Printf("reusing sprite type `%s`\n", name)
	return val, nil
}
