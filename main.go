package main

import (
	"log"

	"github.com/divy-sh/imago/image"
)

func main() {
	img, err := image.Load("testdata/test.png")
	if err != nil {
		log.Fatal(err)
	}
	bright, _ := img.Brighten(50)
	image.SaveImage(bright, "testdata/bright.png")
	value, _ := img.GetGrayScaleByValue()
	image.SaveImage(value, "testdata/grayscale_by_value.png")
	intensity, _ := img.GetGrayScaleByIntensity()
	image.SaveImage(intensity, "testdata/grayscale_by_intensity.jpeg")
	blur, _ := img.Blur()
	image.SaveImage(blur, "testdata/blur.png")
	sharpen, _ := img.Sharpen()
	image.SaveImage(sharpen, "testdata/sharpen.jpg")
	edge, _ := img.EdgeDetect()
	image.SaveImage(edge, "testdata/edge.Png")
	compressed, _ := img.HaarCompress(0.8)
	image.SaveImage(compressed, "testdata/compressed.jpeg")
}
