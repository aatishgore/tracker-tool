package main

import (
	"flag"
	"time"
)

// loading up configuration required
func boostrap() {

	displayWelcomeMessage()

	// parse parameters passed
	parseDebug := flag.Bool("debug", false, "set debug")
	parseLog := flag.Bool("log", true, "set debug")
	flag.Parse()
	debug = *parseDebug
	logInfo = *parseLog

	// check if db file exist, if not create one
	openOrCreateFile("wfh.db", true)
	// setting up first screen shot after 10 sec of app load
	screenShotTrigger = time.Now().Add(time.Second * 10)
	// capture user keyboard press and mouse movement
	go trackPeripheralDevice()
}
