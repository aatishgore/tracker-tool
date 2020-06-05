package main

import (
	"fmt"
	"math/rand"
	"time"
)

func triggerScreenShot(t time.Time) {
	span := agent.ProfileWithName("Capture ScreenShot")
	min := 5
	max := 15
	randNumber := rand.Intn(max-min) + min
	captureScreeShot()
	logger.Printf(" KeyPressed: %v and Mouse moved: %v", keyPress, mouseMovement)

	muTx.Lock()
	keyPress = 0
	mouseMovement = 0
	muTx.Unlock()

	nextTrigger = t.Add(time.Minute * time.Duration(randNumber))
	if true {
		fmt.Printf("\n ScreenShot Captured @ %s \n", t.Format("2006-01-02 15:04:05"))
		fmt.Printf("\n Next Trigger @ %s \n", nextTrigger.Format("2006-01-02 15:04:05"))
	}
	span.Stop()

}
