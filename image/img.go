package image

import (
	"errors"
	"math"
)

type Pixel struct {
	r float64
	g float64
	b float64
	a float64
}

type Img struct {
	p [][]Pixel
	h int
	w int
}

// create a new blank image
func NewImage(height, width int) (*Img, error) {
	if height <= 0 || width <= 0 {
		return nil, errors.New("invalid image dimensions")
	}
	pixels := make([][]Pixel, height)

	for i := range pixels {
		pixels[i] = make([]Pixel, width)
	}
	return &Img{
		p: pixels,
		h: height,
		w: width,
	}, nil
}

// flip image horizontally
func (img *Img) HorizontalFlip() (*Img, error) {
	return process(
		func(i, j int, newImg *Img) {
			newImg.p[i][j] = img.p[i][img.w-j-1]
		}, img,
	)
}

// flip image vertically
func (img *Img) VerticalFlip() (*Img, error) {
	return process(
		func(i, j int, newImg *Img) {
			newImg.p[i][j] = img.p[img.h-i-1][j]
		}, img,
	)
}

func (img *Img) Brighten(bVal float64) (*Img, error) {
	if bVal < -100 || bVal > 100 {
		return nil, errors.New("invalid brightness value")
	}
	increase := bVal / 100
	return process(
		func(i, j int, newImg *Img) {
			newImg.p[i][j].r = clampPixelValue(img.p[i][j].r + increase)
			newImg.p[i][j].g = clampPixelValue(img.p[i][j].g + increase)
			newImg.p[i][j].b = clampPixelValue(img.p[i][j].b + increase)
			newImg.p[i][j].a = img.p[i][j].a
		}, img,
	)
}

func (img *Img) GetRed() (*Img, error) {
	return process(
		func(i, j int, newImg *Img) {
			newImg.p[i][j].r = img.p[i][j].r
			newImg.p[i][j].a = img.p[i][j].a
		}, img,
	)
}

func (img *Img) GetGreen() (*Img, error) {
	return process(
		func(i, j int, newImg *Img) {
			newImg.p[i][j].g = img.p[i][j].g
			newImg.p[i][j].a = img.p[i][j].a
		}, img,
	)
}
func (img *Img) GetBlue() (*Img, error) {
	return process(
		func(i, j int, newImg *Img) {
			newImg.p[i][j].b = img.p[i][j].b
			newImg.p[i][j].a = img.p[i][j].a
		}, img,
	)
}

func (img *Img) GetGrayScaleByValue() (*Img, error) {
	return process(
		func(i, j int, newImg *Img) {
			maxVal := max(max(img.p[i][j].r, img.p[i][j].g), img.p[i][j].b)
			newImg.p[i][j].r = maxVal
			newImg.p[i][j].g = maxVal
			newImg.p[i][j].b = maxVal
			newImg.p[i][j].a = img.p[i][j].a
		}, img,
	)
}

func (img *Img) GetGrayScaleByIntensity() (*Img, error) {
	return process(
		func(i, j int, newImg *Img) {
			average := (img.p[i][j].r + img.p[i][j].g + img.p[i][j].b) / 3
			newImg.p[i][j].r = average
			newImg.p[i][j].g = average
			newImg.p[i][j].b = average
			newImg.p[i][j].a = img.p[i][j].a
		}, img,
	)
}

func (img *Img) HaarCompress(ratio float64) (*Img, error) {
	if ratio < 0 || ratio > 1 {
		return nil, errors.New("invalid compression ratio")
	}
	return haarCompress(img, ratio)
}

func (img *Img) Sharpen() (*Img, error) {
	return sharpenFilter(img)
}

func (img *Img) Blur() (*Img, error) {
	return blurFilter(img)
}

func (img *Img) EdgeDetect() (*Img, error) {
	return edgeDetectFilter(img)
}

func (img *Img) LevelAdjust(blacks, mids, whites float64) (*Img, error) {
	if blacks < 0 || blacks >= mids || mids >= whites || whites > 1 {
		return nil, errors.New("invalid level adjustment values")
	}

	adjust := func(value float64) float64 {
		if value <= blacks {
			return 0
		} else if value >= whites {
			return 1
		} else {
			normalized := (value - blacks) / (whites - blacks)
			gammaCorrected := math.Pow(normalized, math.Log(0.5)/math.Log(mids))
			return gammaCorrected
		}
	}

	return process(
		func(i, j int, newImg *Img) {
			newImg.p[i][j].r = adjust(img.p[i][j].r)
			newImg.p[i][j].g = adjust(img.p[i][j].g)
			newImg.p[i][j].b = adjust(img.p[i][j].b)
			newImg.p[i][j].a = img.p[i][j].a
		}, img,
	)
}

func (img *Img) ColorCorrect() (*Img, error) {
	return colorCorrect(img)
}

func (img *Img) GetHistogram(histSize int) (*Img, error) {
	if histSize < 100 {
		return nil, errors.New("histogram size too small")
	}
	return getHistogram(img, histSize)
}

/*
private functions beyond this point
*/
func clampPixelValue(val float64) float64 {
	if val > 1 {
		val = 1
	} else if val < 0 {
		val = 0
	}
	return val
}

func process(f func(int, int, *Img), img *Img) (*Img, error) {
	newImg, _ := NewImage(img.h, img.w)
	for i := range img.h {
		for j := range img.w {
			f(i, j, newImg)
		}
	}
	return newImg, nil
}
