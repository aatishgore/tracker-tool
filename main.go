package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	"flag"

	"github.com/go-vgo/robotgo"
)

type activeWindow struct {
	activeTime   float64
	activeWindow string
}

var (
	keyPress      int = 0
	mouseMovement int = 0
	nextTrigger   time.Time
	debug         bool
	logInfo       bool
	muTx          = &sync.Mutex{}
	logger        *log.Logger

	prevTitle      string    = ""
	activeWindowOn time.Time = time.Now()
	captureStart   bool      = false
	mpin           string    = ""
)

// this function would execute every specific duration
func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func startWorker(t time.Time) {
	if captureStart {
		fmt.Println("capturing screen")
		var name string = ""
		if runtime.GOOS == "linux" {
			name = runCommand()
		} else {
			// get current active window app name by process id
			// get process id
			pid := robotgo.GetPID()

			// get current active window app name by process id
			name, err := robotgo.FindName(pid)
			if err == nil {
				if name != prevTitle {
					diff := calculateTimeDifference(t, activeWindowOn)
					activeWindowOn = t
					logToDB(prevTitle, diff)
					prevTitle = name
				}
			} else {
				fmt.Println(err)
			}
		}
		fmt.Println(name)

		if t.Sub(nextTrigger) > 0 {
			triggerScreenShot(t)
		}
	} else {
		fmt.Println("not capturing screen")
	}

}
func main() {

	log.SetOutput(ioutil.Discard)
	if runtime.GOOS == "linux" {
		runCommand()
	}
	welcomeMessage()
	//requestMpin()
	// call load up function
	boostrap()

	f, err := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger = log.New(f, "Date", log.LstdFlags)
	nextTrigger = time.Now().Add(time.Second * 10)
	go observerInputMovement()
	// clean function to be called
	setupCloseHandler()
	load()

}

// clean function to be triggred on process quit
func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		now := time.Now()
		diff := calculateTimeDifference(now, activeWindowOn)
		logToDB(prevTitle, diff)
		fmt.Println("\r- Clean up your data")
		copyToLog()
		os.Exit(0)
	}()
}

// loading up configuration required
func boostrap() {
	parseDebug := flag.Bool("debug", false, "set debug")
	parseLog := flag.Bool("log", false, "set debug")
	flag.Parse()
	debug = *parseDebug
	logInfo = *parseLog
	// check if db file exist, if not create one
	touchFile("wfh.db")

}

func welcomeMessage() {
	art :=
		` 	
================================================================================================    
      ___      __   __   _  _   ___    ___  __          ___  __   ___  __   ___ ___  
|  | |__  |   /  ' /  \ | \/ | |__      |  /  \   |\ | |__  /  \ /__  /  \ |__   |   
|/\| |___ |___\__, \__/ |    | |___     |  \__/   | \| |___ \__/ ___/ \__/ |     |   

=================================================================================================  
	 
	
	`
	fmt.Println(art)

}
