package main

import (
	"log"

	"github.com/divy-sh/imago/app"
)

func main() {
	img, err := app.Load("./test.png")
	if err != nil {
		log.Fatal(err)
	}

	img, _ = img.VerticalFlip()
	img.SaveAsPNG("compressed.png")
}
