package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

func captureScreeShot() {
	currentTime := time.Now()
	// Capture each displays.
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		panic("Active display not found")
	}

	time.Sleep(10 * time.Second)

	var all image.Rectangle = image.Rect(0, 0, 0, 0)

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		all = bounds.Union(all)
	}

	// Capture all desktop region into an image.
	img, err := screenshot.Capture(all.Min.X, all.Min.Y, all.Dx(), all.Dy())
	if err != nil {
		panic(err)
	}
	fileName := fmt.Sprintf("all_%s.png", currentTime.Format("2006-01-02 15:04:05"))

	save(img, fileName)
	//file upload to server
	// path, err := os.Getwd()
	// if err != nil {
	// 	log.Println(err)
	// }

	// filePath := filepath.Join(path, fileName)
	// uploadFile(filePath)
}

// save *image.RGBA to filePath with PNG format.
func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}
