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

	val, _ := img.GetGrayScaleByValue()
	val.SaveAsPNG("testdata/grayscale_by_value.png")
	intensity, _ := img.GetGrayScaleByIntensity()
	intensity.SaveAsPNG("testdata/grayscale_by_intensity.png")

}
