package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/go-vgo/robotgo"
)

// trackPeripheralDevice is a function which capture every key press and mouse movement of user desktop
func trackPeripheralDevice() {

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

// this function stores active window information like how much time was spend on specific window
func storeCurrentActiveWindowName(t time.Time) {
	var (
		activeWindowName string = ""
		err              error
	)

	if runtime.GOOS == "linux" {
		activeWindowName = getLinuxActiveWindowName()
	} else {
		// get current active window app process id
		pid := robotgo.GetPID()

		// get current active window app name by process id
		activeWindowName, err = robotgo.FindName(pid)
		if err != nil {
			activeWindowName = ""
		}
	}
	if activeWindowName != prevTitle {
		diff := calculateTimeDifference(t, activeWindowOn)
		logUserActivityInDB(prevTitle, diff)
		activeWindowOn = t
		prevTitle = activeWindowName
	}

	if debug {
		fmt.Println(activeWindowName)
	}

}
