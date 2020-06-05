package main

import (
	"math"

	"github.com/go-vgo/robotgo"
)

func observerInputMovment() {

	evChan := robotgo.EventStart()
	for e := range evChan {
		// Key press event
		if e.Kind == 3 {
			keyPress++
		}
		if e.Kind == 9 {
			x := e.X
			y := e.Y
			if prevX != x && prevY != y {

				var distance float64 = 0
				xSquared := math.Pow(float64(prevX-x), 2)
				ySquared := math.Pow(float64(prevY-y), 2)
				distance = math.Sqrt(float64(xSquared + ySquared))

				if int(distance) > 10 {
					mouseMovement++
				}
				prevX = x
				prevY = y
			}
		}
	}
}
