package main

import (
	"fmt"
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

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
	erodeTemplate := gocv.GetStructuringElement(gocv.MorphRect, image.Point{X: 3, Y: 3})
	dilateTemplate := gocv.GetStructuringElement(gocv.MorphRect, image.Point{X: 35, Y: 35})

	// create variable image
	gray := gocv.NewMat()
	result := gocv.NewMat()
	before := gocv.NewMat()
	after := gocv.NewMat()

	// indicator when video just open
	var isFirst bool = true

	var human, motorcycle int

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

		// copy after to before
		before = after.Clone()

		gocv.Threshold(result, &result, 45, 255, gocv.ThresholdBinary)
		gocv.Erode(result, &result, erodeTemplate)
		gocv.Dilate(result, &result, dilateTemplate)

		moments := gocv.Moments(result, true)
		point := image.Point{
			X: int(moments["m10"] / moments["m00"]),
			Y: int(moments["m01"] / moments["m00"]),
		}

		// create dot (center of object)
		gocv.Circle(&img, point, 5, color.RGBA{R: 255}, -1)

		// count blob and save it to variable
		contours := gocv.FindContours(result, gocv.RetrievalList, gocv.ChainApproxSimple)
		var pixels []int
		for _, contour := range contours {
			if len(contour) > 75 {
				pixels = append(pixels, len(contour))
			}
		}

		for _, pixel := range pixels {
			// check the blob is already passed centerline or not yet
			if point.X > 293 && point.X < 310 {
				// check blob is motorcycle or human
				if pixel > 75 && pixel <= 110 {
					human++
				} else if pixel > 110 && pixel <= 150 {
					motorcycle++
				}
			}
		}

		fmt.Println("human:", human, "motorcycle:", motorcycle)

		originalWindow.IMShow(img)
		originalWindow.WaitKey(1) // set 50 to normal video
	}
}
