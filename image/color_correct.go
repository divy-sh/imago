package image

import "fmt"

// TODO - fix the color correction algorithm
func colorCorrect(img *Img) (*Img, error) {
	histSize := 256
	redP, greenP, blueP := getHistogramData(img, histSize)
	red, green, blue := *redP, *greenP, *blueP
	rAvg, gAvg, bAvg := getAvgHistColorValues(&red, &green, &blue, histSize)
	average := (rAvg + gAvg + bAvg) / 3
	rShift := rAvg - average
	gShift := gAvg - average
	bShift := bAvg - average
	fmt.Println(rShift, gShift, bShift)
	return process(
		func(i, j int, newImg *Img) {
			newImg.p[i][j].r = clampPixelValue(img.p[i][j].r + rShift)
			newImg.p[i][j].g = clampPixelValue(img.p[i][j].g + gShift)
			newImg.p[i][j].b = clampPixelValue(img.p[i][j].b + bShift)
			newImg.p[i][j].a = img.p[i][j].a
		}, img,
	)
}

func getAvgHistColorValues(red, green, blue *[]int, histSize int) (float64, float64, float64) {
	rAvg, gAvg, bAvg := 0, 0, 0

	for i := 0; i < histSize; i++ {
		rAvg += (*red)[i]
		gAvg += (*green)[i]
		bAvg += (*blue)[i]
	}
	return float64(rAvg) / float64(histSize), float64(gAvg) / float64(histSize), float64(bAvg) / float64(histSize)
}
