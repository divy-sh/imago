package image

func getHistogram(img *Img, histSize int) (*Img, error) {
	hist, _ := NewImage(histSize+1, histSize+1)

	for i := 0; i < histSize; i++ {
		for j := 0; j < histSize; j++ {
			hist.p[i][j].a = 1
		}
	}

	redP, greenP, blueP := getHistogramData(img, histSize)
	red, green, blue := *redP, *greenP, *blueP

	maxVal := 0
	for i := 0; i < histSize; i++ {
		maxVal = max(maxVal, red[i])
		maxVal = max(maxVal, green[i])
		maxVal = max(maxVal, blue[i])
	}

	return drawHistogram(hist, red, green, blue, maxVal, histSize)
}

func getHistogramData(img *Img, histSize int) (*[]int, *[]int, *[]int) {
	red := make([]int, histSize+1)
	green := make([]int, histSize+1)
	blue := make([]int, histSize+1)

	for i := 0; i < img.h; i++ {
		for j := 0; j < img.w; j++ {
			red[int(img.p[i][j].r*float64(histSize))]++
			green[int(img.p[i][j].g*float64(histSize))]++
			blue[int(img.p[i][j].b*float64(histSize))]++

		}
	}
	return &red, &green, &blue
}

func drawHistogram(hist *Img, red, green, blue []int, maxVal, histSize int) (*Img, error) {
	for x := 0; x < histSize; x++ {
		rVal := red[x] * histSize / maxVal
		gVal := green[x] * histSize / maxVal
		bVal := blue[x] * histSize / maxVal
		for y := histSize; y >= histSize-rVal; y-- {
			hist.p[y][x].r = 1
		}
		for y := histSize; y >= histSize-bVal; y-- {
			hist.p[y][x].g = 1
		}
		for y := histSize; y >= histSize-gVal; y-- {
			hist.p[y][x].b = 1
		}
	}

	return hist, nil
}
