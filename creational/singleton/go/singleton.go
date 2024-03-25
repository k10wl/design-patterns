package main

import (
	"fmt"
	"sync"
)

type Singleton struct{}

var (
	mutex    = &sync.Mutex{}
	instance *Singleton
)

func getSingleton() *Singleton {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			fmt.Println("does not exist, creating new one")
			instance = &Singleton{}
		} else {
			fmt.Println("exists on other thread")
		}
	} else {
		fmt.Println("exists on this thread")
	}
	return instance
}

func (s *Singleton) Operation() {
	fmt.Println("performing operation")
}
