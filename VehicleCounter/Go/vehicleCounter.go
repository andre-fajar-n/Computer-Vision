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

	// create template for erode and dilate
	// erodeTemplate := gocv.GetStructuringElement(gocv.MorphRect, image.Point{X: -1, Y: -1})

	// create variable image
	gray := gocv.NewMat()
	result := gocv.NewMat()
	before := gocv.NewMat()
	after := gocv.NewMat()

	// indicator when video just open
	var isFirst bool = true

	// show the video frame by frame
	for {
		// to break terminal when video is end
		if ok := video.Read(&img); !ok || img.Empty() {
			break
		}

		// convert to gray image
		gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

		// save image before and now
		after = gray.Clone()
		if isFirst {
			before = after.Clone()
			result = after.Clone()
			isFirst = false
		}

		// subtract before and after
		gocv.AbsDiff(after, before, &result)

		//
		before = after.Clone()
		gocv.Threshold(result, &result, 45, 255, gocv.ThresholdBinary)

		originalWindow.IMShow(result)
		originalWindow.WaitKey(1)
	}
}
