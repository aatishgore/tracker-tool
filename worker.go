package main

import (
	"fmt"
	"time"
)

func startWorker(t time.Time) {
	if trackingStart {
		timeStampString := t.Format("2006-01-02 15:04:05")
		layOut := "2006-01-02 15:04:05"
		timeStamp, _ := time.Parse(layOut, timeStampString)
		hr, min, sec := timeStamp.Clock()
		fmt.Println(hr, min, sec)
		if min%10 == 0 && sec == 0 {
			saveLogs()
		}

		if debug {
			fmt.Println("capturing screen @", t.Format("2006-01-02 15:04:05"))
		}
		storeCurrentActiveWindowName(t)
		if t.Sub(screenShotTrigger) > 0 {
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
