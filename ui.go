package main

import (
	"fmt"
	"os"

	"github.com/martinlindhe/inputbox"
)

func requestMpin() {
	got, ok := inputbox.InputBox("Dialog title", "Type a number", "abc")
	if ok {
		mpin = got
	} else {
		fmt.Println("Please enter Mpin number")
		os.Exit(0)

	}
}
