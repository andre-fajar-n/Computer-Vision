package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

// func findBlob(img Mat) int64 {

// }

func main() {
	// declare window
	originalWindow := gocv.NewWindow("original")

	// get video file
	videoPath := "../data/MV_v2.avi"
	video, err := gocv.VideoCaptureFile(videoPath)
	if err != nil {
		fmt.Printf("Cannot open the video")
	}
	defer video.Close()

	// create variable image
	img := gocv.NewMat()
	defer img.Close()

	// show the video frame by frame
	for {
		// to break terminal when video is end
		if ok := video.Read(&img); !ok || img.Empty() {
			break
		}

		originalWindow.IMShow(img)
		originalWindow.WaitKey(50)
	}
}
