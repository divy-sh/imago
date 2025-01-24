package image

func blurFilter(img *Img) (*Img, error) {
	blurFilter := [][]float64{
		{1.0 / 16, 1.0 / 8, 1.0 / 16},
		{1.0 / 8, 1.0 / 4, 1.0 / 8},
		{1.0 / 16, 1.0 / 8, 1.0 / 16},
	}
	return applyFilter(img, &blurFilter)
}

func sharpenFilter(img *Img) (*Img, error) {
	sharpenFilter := [][]float64{
		{-1.0 / 8, -1.0 / 8, -1.0 / 8, -1.0 / 8, -1.0 / 8},
		{-1.0 / 8, 1.0 / 4, 1.0 / 4, 1.0 / 4, -1.0 / 8},
		{-1.0 / 8, 1.0 / 4, 1.0, 1.0 / 4, -1.0 / 8},
		{-1.0 / 8, 1.0 / 4, 1.0 / 4, 1.0 / 4, -1.0 / 8},
		{-1.0 / 8, -1.0 / 8, -1.0 / 8, -1.0 / 8, -1.0 / 8},
	}
	return applyFilter(img, &sharpenFilter)
}

func edgeDetectFilter(img *Img) (*Img, error) {
	edgeFilter := [][]float64{
		{-1.0, -1.0, -1.0},
		{-1.0, 8.0, -1.0},
		{-1.0, -1.0, -1.0},
	}
	return applyFilter(img, &edgeFilter)
}

/*
Private functions below this point
*/

func applyFilter(img *Img, filterPointer *[][]float64) (*Img, error) {
	filter := *filterPointer
	return process(
		func(h, w int, newImg *Img) {
			steps := len(filter) / 2
			r := 0.
			g := 0.
			b := 0.
			for i := -steps; i <= steps; i++ {
				for j := -steps; j <= steps; j++ {
					if h+i >= 0 && h+i < img.h && w+j >= 0 && w+j < img.w {
						r += filter[i+steps][j+steps] * img.p[h+i][w+j].r
						g += filter[i+steps][j+steps] * img.p[h+i][w+j].g
						b += filter[i+steps][j+steps] * img.p[h+i][w+j].b
					}
				}
			}
			newImg.p[h][w].r = clampPixelValue(r)
			newImg.p[h][w].g = clampPixelValue(g)
			newImg.p[h][w].b = clampPixelValue(b)
			newImg.p[h][w].a = img.p[h][w].a
		}, img,
	)
}
