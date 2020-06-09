package main

import (
	"github.com/go-vgo/robotgo"
)

func observerInputMovement() {

	evChan := robotgo.EventStart()
	for e := range evChan {
		// Key press event
		if e.Kind == 3 {
			keyPress++
		}
		// Mouse movement event
		if e.Kind == 9 {
			mouseMovement++
		}
	}
}
