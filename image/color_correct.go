package image

import "fmt"

// TODO - fix the color correction algorithm
func colorCorrect(img *Img) (*Img, error) {
	histSize := 256
	redP, greenP, blueP := getHistogramData(img, histSize)
	red, green, blue := *redP, *greenP, *blueP
	rPeak, gPeak, bPeak := getMaxHistColorValues(&red, &green, &blue, histSize)
	average := (rPeak + gPeak + bPeak) / 3
	return process(
		func(i, j int, newImg *Img) {
			newImg.p[i][j].r = clampPixelValue(img.p[i][j].r + (float64(average-rPeak))/float64(histSize))
			newImg.p[i][j].g = clampPixelValue(img.p[i][j].g + (float64(average-gPeak))/float64(histSize))
			newImg.p[i][j].b = clampPixelValue(img.p[i][j].b + (float64(average-bPeak))/float64(histSize))
			fmt.Println(newImg.p[i][j])
			newImg.p[i][j].a = img.p[i][j].a
		}, img,
	)
}

func getMaxHistColorValues(red, green, blue *[]int, histSize int) (int, int, int) {
	rMax, gMax, bMax := 0, 0, 0

	for i := 50; i < histSize-50; i++ {
		rMax = max(rMax, (*red)[i])
		gMax = max(gMax, (*green)[i])
		bMax = max(bMax, (*blue)[i])
	}
	return rMax, gMax, bMax
}
