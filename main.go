package main

import (
	"log"

	"github.com/divy-sh/imago/app"
)

func main() {
	img, err := app.Load("testdata/test.png")
	if err != nil {
		log.Fatal(err)
	}
	bright, _ := img.Brighten(50)
	app.SaveImage(bright, "testdata/bright")
	value, _ := img.GetGrayScaleByValue()
	app.SaveImage(value, "testdata/grayscale_by_value")
	intensity, _ := img.GetGrayScaleByIntensity()
	app.SaveImage(intensity, "testdata/grayscale_by_intensity")
	blur, _ := img.Blur()
	app.SaveImage(blur, "testdata/blur")
	sharpen, _ := img.Sharpen()
	app.SaveImage(sharpen, "testdata/sharpen")
	edge, _ := img.EdgeDetect()
	app.SaveImage(edge, "testdata/edge.Png")
	compressed, _ := img.HaarCompress(0.8)
	app.SaveImage(compressed, "testdata/compressed")
}
