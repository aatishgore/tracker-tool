package main

import (
	"fmt"
	"log"
	"time"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

var w *astilectron.Window

func load() {
	// Set logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	// Create astilectron
	a, err := astilectron.New(l, astilectron.Options{
		AppName:           "Test",
		BaseDirectoryPath: "example",
	})
	if err != nil {
		l.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}
	defer a.Close()

	// Handle signals
	a.HandleSignals()

	// Start
	if err = a.Start(); err != nil {
		l.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	}

	if w, err = a.NewWindow("http://localhost:8000/", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(700),
		Width:  astikit.IntPtr(700),
	}); err != nil {
		l.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}

	// Create windows
	if err = w.Create(); err != nil {
		l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}
	w.SendMessage("hello", func(m *astilectron.EventMessage) {
		// Unmarshal

		// Process message
		log.Printf("received\n")
	})
	w.OnMessage(func(m *astilectron.EventMessage) interface{} {
		// Unmarshal
		var s string
		m.Unmarshal(&s)

		// Process message

		switch s {
		case "start":
			captureStart = true
			break
		case "stop":
			captureStart = false
			break
		default:
			fmt.Println("unassigned message")
		}

		return nil
	})
	w.OpenDevTools()
	// Blocking pattern

	doEvery(1*time.Second, startWorker)

	a.Wait()
}

func sendData(message string) {
	w.SendMessage(message, func(m *astilectron.EventMessage) {
		// Unmarshal

		// Process message
		log.Printf("received\n")
	})
	fmt.Print(&w)

}
