package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/stackimpact/stackimpact-go"
)

type activeWindow struct {
	activeTime   float64
	activeWindow string
}

var (
	prevX          int16 = 0
	prevY          int16 = 0
	keyPress       int   = 0
	mouseMovement  int   = 0
	nextTrigger    time.Time
	debug          bool = false
	muTx                = &sync.Mutex{}
	logger         *log.Logger
	agent          *stackimpact.Agent
	prevTitle      string    = ""
	activeWindowOn time.Time = time.Now()
)
var activeWindows = make([]activeWindow, 0)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func startWorker(t time.Time) {
	if debug {
		fmt.Printf("Current Time is %s \nTrigger Time is %s \n condition is %v\n ", t.Format("2006-01-02 15:04:05"), nextTrigger.Format("2006-01-02 15:04:05"), t.Sub(nextTrigger))
	}
	pid := robotgo.GetPID()
	title := robotgo.GetTitle()

	name, err := robotgo.FindName(pid)
	if err == nil {
		if name != prevTitle {
			diff := int(t.Sub(activeWindowOn).Seconds())
			activeWindowOn = t
			prevTitle = name
			logToDB(prevTitle, diff)
			// activeWindows = append(activeWindows, activeWindow{
			// 	diff,
			// 	prevTitle,
			// })
		}
	} else {
		fmt.Println(err, title)
	}

	if t.Sub(nextTrigger) > 0 {
		triggerScreenShot(t)
	}

}
func main() {

	agent = stackimpact.Start(stackimpact.Options{
		AgentKey: "924ed3987351cf81f7af8b9431eff889720c6a13",
		AppName:  "MyGoApp",
	})

	f, err := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger = log.New(f, "Date", log.LstdFlags)
	// nextTrigger = time.Now().Add(time.Second * 10)
	go observerInputMovment()
	setupCloseHandler()
	doEvery(1*time.Second, startWorker)

}

func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Clean up your data")
		copyToLog()
		os.Exit(0)
	}()
}
