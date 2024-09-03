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

	img, _ = img.GetGrayScaleByValue()
	img.SaveAsPNG("testdata/grayscale_by_value.png")
}
