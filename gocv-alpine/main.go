package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gocv.io/x/gocv"
)

func main() {
	// argument
	saveDir := flag.String("save", "/out", "Save Image Directory")
	rtspURL := flag.String("rtsp", "rtsp://admin:123456@192.168.0.55:554/0/video1", "RTSP URL")
	flag.Parse()

	fmt.Println("Show RTSP URL:", *rtspURL)
	fmt.Println("Save to :", *saveDir)

	// check saveDir exist ,if not create it
	if _, err := os.Stat(*saveDir); os.IsNotExist(err) {
		os.MkdirAll(*saveDir, os.ModePerm)
	}

	// FPS
	fps := 0
	go func() {
		c := time.Tick(1 * time.Second)
		for now := range c {
			fmt.Printf("%v  FPS: %d\n", now, fps)
			fps = 0
		}
	}()

	// show and save
	rtspSource, _ := gocv.OpenVideoCapture(*rtspURL)
	img := gocv.NewMat()

	for {
		rtspSource.Read(&img)
		// millisecond is enough for image input
		path := filepath.Join(*saveDir, strconv.FormatInt(time.Now().UnixNano()/1000, 10)+".jpg")
		gocv.IMWrite(path, img)
		fmt.Println("Save to", path)
		fps++
	}
}
