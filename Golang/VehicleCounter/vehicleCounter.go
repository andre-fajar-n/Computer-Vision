package main

import (
	"fmt"
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

var (
	gray       = gocv.NewMat()
	result     = gocv.NewMat()
	before     = gocv.NewMat()
	after      = gocv.NewMat()
	isFirst    = true
	human      = 0
	motorcycle = 0
)

func main() {
	// declare window
	// originalWindow := gocv.NewWindow("original")

	red := color.RGBA{R: 255}

	// get video file
	videoPath := "../../data/vehicleCounter.avi"
	video, err := gocv.VideoCaptureFile(videoPath)
	if err != nil {
		fmt.Println("Cannot open the video")
	}
	defer video.Close()

	// create variable image
	img := gocv.NewMat()
	defer img.Close()

	// create template for erode and dilate
	erodeTemplate := gocv.GetStructuringElement(gocv.MorphRect, image.Point{X: 3, Y: 3})
	dilateTemplate := gocv.GetStructuringElement(gocv.MorphRect, image.Point{X: 35, Y: 35})

	if ok := video.Read(&img); !ok {
		fmt.Println("cannot read file")
	}
	// preparation to save result video
	saveFile := "../../result/result.avi"
	writer, err := gocv.VideoWriterFile(saveFile, "MJPG", 50, img.Cols()*2, img.Rows()*2, true)
	if err != nil {
		fmt.Printf("error opening video writer device: %v\n", err)
		return
	}
	defer writer.Close()

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
		gocv.Circle(&img, point, 5, red, -1)

		// count blob and save it to variable
		contours := gocv.FindContours(result, gocv.RetrievalList, gocv.ChainApproxSimple)
		var pixels []int
		for idx, contour := range contours {
			if len(contour) > 75 {
				pixels = append(pixels, len(contour))
				gocv.DrawContours(&img, contours, idx, red, 2)
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

		text := fmt.Sprintf("human: %d, motorcycle: %d", human, motorcycle)

		// add centerline
		centerLine := img.Cols() / 2
		gocv.Line(&img, image.Point{X: centerLine}, image.Point{X: centerLine, Y: img.Rows()}, red, 2)

		// add text
		fmt.Println(text)
		gocv.PutText(&img, text, image.Point{X: 10, Y: img.Rows() - 50}, gocv.FontHersheySimplex, 1.1, red, 3)

		writer.Write(img)
		// originalWindow.IMShow(img)
		// originalWindow.WaitKey(50) // set 50 to normal video
	}
}
