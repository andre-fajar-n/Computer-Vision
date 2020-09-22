package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

var redColor = color.RGBA{R: 255}
var posX, posY int

const minPos = 9.9e+25

func main() {
	// declare window
	window1 := gocv.NewWindow("image 1")
	window2 := gocv.NewWindow("image 2")

	imgPath1 := "./data/image1-a.jpg"
	imgPath2 := "./data/image1-b.jpg"
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

	segment := 4

	col := img1.Cols() / segment
	row := img1.Rows() / segment

	for i := 0; i < segment; i++ {
		for j := 0; j < segment; j++ {
			// crop image
			tmp := img1.Region(image.Rect(i*col, j*row, (i+1)*col, (j+1)*row))

			result := img1.Clone()

			// match template
			mask := gocv.NewMat()
			gocv.MatchTemplate(img2, tmp, &result, gocv.TmSqdiff, mask)
			mask.Close()
			minVal, maxVal, minLoc, maxLoc := gocv.MinMaxLoc(result)

			fmt.Println(minVal, maxVal)
			fmt.Println(minLoc.X, minLoc.Y, maxLoc.X, maxLoc.Y)
			fmt.Println(i*col, j*row, (i+1)*col, (j+1)*row)
			fmt.Println("+++++++++++++++++++++++")

			if minVal < minPos {
				posX = minLoc.X
				// rect := image.Rect(i*col, j*row, (i+1)*col, (j+1)*row)
				// gocv.Rectangle(&img1, rect, redColor, 1)
			}
		}
	}
	fmt.Println(posX)
	for {
		window1.IMShow(img1)
		window2.IMShow(img2)
		window1.WaitKey(1)
	}
}
