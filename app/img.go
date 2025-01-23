package app

import (
	"errors"
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

// TODO implement this
func (img *Img) LevelAdjust(blacks, mids, whites float64) (*Img, error) {
	return nil, nil
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
