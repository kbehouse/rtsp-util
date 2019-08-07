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
	rtstURL := os.Args[1]

	fmt.Println("Show URL:", rtstURL)

	fps:=0
	go func() {
		c := time.Tick(1 * time.Second)
		for now := range c {
			fmt.Printf("%v  FPS: %d\n", now, fps)
			fps = 0
		}
	}()

	webcam, _ := gocv.OpenVideoCapture(rtstURL)
	window := gocv.NewWindow(rtstURL)
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		// fmt.Printf("time: %v, size: %v\n", time.Now(), img.Size())
		window.IMShow(img)
		window.WaitKey(1)
		fps++
	}
}
