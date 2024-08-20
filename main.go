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

	img, _ = img.Compress(.5)
	img.SaveAsPNG("compressed.png")
}
