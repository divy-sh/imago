package main

import (
	"log"

	"github.com/divy-sh/imago/app"
)

func main() {
	img, err := app.Load("testdata/test.Png")
	if err != nil {
		log.Fatal(err)
	}
	bright, _ := img.Brighten(50)
	app.SaveAsPng(bright, "testdata/bright.Png")
	value, _ := img.GetGrayScaleByValue()
	app.SaveAsPng(value, "testdata/grayscale_by_value.Png")
	intensity, _ := img.GetGrayScaleByIntensity()
	app.SaveAsPng(intensity, "testdata/grayscale_by_intensity.Png")
	blur, _ := img.Blur()
	app.SaveAsPng(blur, "testdata/blur.Png")
	sharpen, _ := img.Sharpen()
	app.SaveAsPng(sharpen, "testdata/sharpen.Png")
	edge, _ := img.EdgeDetect()
	app.SaveAsPng(edge, "testdata/edge.Png")
	compressed, _ := img.HaarCompress(0.8)
	app.SaveAsPng(compressed, "testdata/compressed.Png")
}
