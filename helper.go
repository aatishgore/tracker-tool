package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func triggerScreenShot(t time.Time) {

	captureStatus := captureScreeShot()

	if !captureStatus {
		logger.Printf("No active screen found")
	}
	msg := fmt.Sprintf(" KeyPressed: %v and Mouse moved: %v", keyPress, mouseMovement)
	if displayUI {
		sendData(msg)
	}

	if logInfo {
		logger.Printf(" KeyPressed: %v and Mouse moved: %v", keyPress, mouseMovement)
	}
	muTx.Lock()
	keyPress = 0
	mouseMovement = 0
	muTx.Unlock()

	randNumber := rand.Intn(maxWaitTimeForScreenShot-minWaitTimeForScreenShot) + minWaitTimeForScreenShot
	nextTrigger = t.Add(time.Minute * time.Duration(randNumber))
	if debug {
		fmt.Printf("\n ScreenShot Captured @ %s \n", t.Format("2006-01-02 15:04:05"))
		fmt.Printf("\n Next Trigger @ %s \n", nextTrigger.Format("2006-01-02 15:04:05"))
	}
}
func openOrCreateFile(name string, closeFile bool) (*os.File, error) {
	file, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	if closeFile {
		defer file.Close()
	}
	return file, nil
}

func calculateTimeDifference(t time.Time, compare time.Time) int {
	diff := int(t.Sub(compare).Seconds())
	return diff
}

// displayWelcomeMessage is a function which display welcome message on app boot
func displayWelcomeMessage() {
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
