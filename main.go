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
	bright.SaveAsPNG("testdata/bright.png")
	value, _ := img.GetGrayScaleByValue()
	value.SaveAsPNG("testdata/grayscale_by_value.png")
	intensity, _ := img.GetGrayScaleByIntensity()
	intensity.SaveAsPNG("testdata/grayscale_by_intensity.png")
	blur, _ := img.Blur()
	blur.SaveAsPNG("testdata/blur.png")
	sharpen, _ := img.Sharpen()
	sharpen.SaveAsPNG("testdata/sharpen.png")
}
