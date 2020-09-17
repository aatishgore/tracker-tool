package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

type activeWindow struct {
	activeTime   float64
	activeWindow string
}

type windowLog struct {
	AppName  string `json:"appName"`
	Duration string `json:"duration"`
}

type logRequest struct {
	AppLogs  []windowLog `json:"app_logs"`
	Mouse    int         `json:"mouse"`
	Keyboard int         `json:"keyboard"`
}

type encryptedLogRequest struct {
	Data string `json:"data"`
}

type imageRequest struct {
	Base64Image string `json:"base64_image"`
}

var (
	keyPress          int
	mouseMovement     int
	screenShotTrigger time.Time
	debug             bool
	logInfo           bool
	muTx              = &sync.Mutex{}
	logger            *log.Logger

	prevTitle                string
	activeWindowOn           time.Time = time.Now()
	trackingStart            bool      = false
	mpin                     string
	minWaitTimeForScreenShot int = 5
	maxWaitTimeForScreenShot int = 10
	gsConnWS                 []*websocket.Conn
	gsMessageType            int
	token                    string
	encryptionKey            string = "n30WF|-|"
)

type socketMessage struct {
	Channel string
	Message string
}

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
	go socketInit()
	doEvery(1*time.Second, startWorker)

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
		os.Exit(0)
	}()
}
