package main

import (
	"fmt"
	"time"
)

func startWorker(t time.Time) {

	if trackingStart {
		if debug {
			fmt.Println("capturing screen @", t.Format("2006-01-02 15:04:05"))
		}
		storeCurrentActiveWindowName(t)
		if t.Sub(nextTrigger) > 0 {
			triggerScreenShot(t)
		}
	}

}

// this function would execute every specific duration
func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}
