package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type activeWindow struct {
	activeTime   float64
	activeWindow string
}

var (
	keyPress      int
	mouseMovement int
	nextTrigger   time.Time
	debug         bool
	logInfo       bool
	displayUI     bool
	muTx          = &sync.Mutex{}
	logger        *log.Logger

	prevTitle                string
	activeWindowOn           time.Time = time.Now()
	trackingStart            bool      = false
	mpin                     string
	minWaitTimeForScreenShot int = 5
	maxWaitTimeForScreenShot int = 15
)

func main() {
	// disable log of electron js app
	//log.SetOutput(ioutil.Discard)

	// call load up function
	boostrap()
	// handle log file
	if logInfo {
		f, err := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		logger = log.New(f, "Date", log.LstdFlags)

	}
	// clean function to be called
	setupCloseHandler()

	if displayUI {
		load()
	} else {
		doEvery(1*time.Second, startWorker)

	}

}

// clean function to be triggred on process quit
func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		now := time.Now()
		diff := calculateTimeDifference(now, activeWindowOn)
		logUserActivityInDB(prevTitle, diff)
		fmt.Println("\r- Clean up your data")
		// TODO: Change this api call on exist
		copyToLog()
		os.Exit(0)
	}()
}
