package main

import (
	"fmt"
	"os"
	"time"

	"gocv.io/x/gocv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("You should inpout RTSP URL, len(os.Args) != 2")
		return
	}
	rtspURL := os.Args[1]

	fmt.Println("Show URL:", rtspURL)

	fps := 0
	go func() {
		c := time.Tick(1 * time.Second)
		for now := range c {
			fmt.Printf("%v  FPS: %d\n", now, fps)
			fps = 0
		}
	}()

	webcam, _ := gocv.OpenVideoCapture(rtspURL)
	window := gocv.NewWindow(rtspURL)
	window.ResizeWindow(800, 600)
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		fmt.Printf("time: %v, size: %v\n", time.Now(), img.Size())
		window.IMShow(img)
		window.WaitKey(1)
		fps++
	}
}
