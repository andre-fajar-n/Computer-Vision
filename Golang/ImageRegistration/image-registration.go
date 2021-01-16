package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

var redColor = color.RGBA{R: 255}
var bx, by, ax, ay int
var segment int = 5

const minPos = 9.9e+25

func main() {
	// declare window
	// window1 := gocv.NewWindow("image 1")
	// window2 := gocv.NewWindow("image 2")
	// windowResult := gocv.NewWindow("result")

	imgPath1 := "../../data/imageRegistration-1.jpg"
	imgPath2 := "../../data/imageRegistration-2.jpg"
	img1 := gocv.IMRead(imgPath1, gocv.IMReadColor)
	img2 := gocv.IMRead(imgPath2, gocv.IMReadColor)
	if img1.Empty() {
		fmt.Printf("Could not read image %s\n", imgPath1)
		os.Exit(1)
	}
	if img2.Empty() {
		fmt.Printf("Could not read image %s\n", imgPath2)
		os.Exit(1)
	}

	col := img1.Cols() / segment
	row := img1.Rows() / segment

	for i := 0; i < segment; i++ {
		for j := 0; j < segment; j++ {
			// create rectangle
			rect := image.Rect(i*col, j*row, (i+1)*col, (j+1)*row)

			// crop image
			tmp := img1.Region(rect)

			result := img1.Clone()

			// match template
			mask := gocv.NewMat()
			gocv.MatchTemplate(img2, tmp, &result, gocv.TmSqdiff, mask)
			mask.Close()
			minVal, _, minLoc, _ := gocv.MinMaxLoc(result)

			if minVal < minPos {
				bx = minLoc.X
				by = minLoc.Y
				ax = i * col
				ay = j * row
			}
		}
	}

	gocv.Rectangle(&img1, image.Rect(ax, ay, ax+col, ay+row), redColor, 1)
	gocv.Rectangle(&img2, image.Rect(bx, by, bx+col, by+row), redColor, 1)
	result := gocv.NewMatWithSize(ay+img2.Rows()-by, ax+img2.Cols()-bx, gocv.MatTypeCV8UC3)

	// attach image1 to result
	roi := image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: img1.Size()[1], Y: img1.Size()[0]},
	}
	resultRoi := result.Region(roi)
	gocv.Resize(img1, &resultRoi, roi.Size(), 0, 0, gocv.InterpolationLinear)

	// attach image2 to result
	roi = image.Rectangle{
		Min: image.Point{X: ax - bx, Y: ay - by},
		Max: image.Point{X: result.Size()[1], Y: result.Size()[0]},
	}
	defer result.Close()
	resultRoi = result.Region(roi)
	defer resultRoi.Close()
	gocv.Resize(img2, &resultRoi, roi.Size(), 0, 0, gocv.InterpolationLinear)

	gocv.IMWrite("../../result/imageRegistrationResult.png", result)

	// window1.IMShow(img1)
	// window2.IMShow(img2)
	// windowResult.IMShow(result)
	// windowResult.WaitKey(0)
}
