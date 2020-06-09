package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func triggerScreenShot(t time.Time) {
	min := 5
	max := 15
	randNumber := rand.Intn(max-min) + min

	captureScreeShot()

	msg := fmt.Sprintf(" KeyPressed: %v and Mouse moved: %v", keyPress, mouseMovement)
	sendData(msg)

	if logInfo {
		logger.Printf(" KeyPressed: %v and Mouse moved: %v", keyPress, mouseMovement)
	}
	muTx.Lock()
	keyPress = 0
	mouseMovement = 0
	muTx.Unlock()

	nextTrigger = t.Add(time.Minute * time.Duration(randNumber))
	if debug {
		fmt.Printf("\n ScreenShot Captured @ %s \n", t.Format("2006-01-02 15:04:05"))
		fmt.Printf("\n Next Trigger @ %s \n", nextTrigger.Format("2006-01-02 15:04:05"))
	}
}
func touchFile(name string) error {
	file, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	return file.Close()
}

func calculateTimeDifference(t time.Time, compare time.Time) int {
	diff := int(t.Sub(compare).Seconds())
	return diff
}
