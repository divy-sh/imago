package image

import (
	"errors"
	"math"
	"sort"
)

// haarCompress compresses an image using the Haar wavelet transform
func haarCompress(img *Img, ratio float64) (*Img, error) {
	if ratio < 0 || ratio > 1 {
		return nil, errors.New("invalid compression ratio")
	}

	size := max(nextPerfectPowerOf2(img.h), nextPerfectPowerOf2(img.w))
	imgData := make([][][]float64, 3)
	for i := range imgData {
		imgData[i] = make([][]float64, size)
		for j := range imgData[i] {
			imgData[i][j] = make([]float64, size)
		}
	}

	for i := 0; i < img.h; i++ {
		for j := 0; j < img.w; j++ {
			imgData[0][i][j] = img.p[i][j].r
			imgData[1][i][j] = img.p[i][j].g
			imgData[2][i][j] = img.p[i][j].b
		}
	}

	for i := 0; i < 3; i++ {
		haarTransform2d(imgData[i], size)
		setValuesZero(imgData[i], ratio)
		inverseHaarTransform2d(imgData[i], size)
	}

	return process(
		func(i, j int, newImg *Img) {
			newImg.p[i][j].r = math.Min(float64(img.h), math.Max(0, imgData[0][i][j]))
			newImg.p[i][j].g = math.Min(float64(img.h), math.Max(0, imgData[1][i][j]))
			newImg.p[i][j].b = math.Min(float64(img.h), math.Max(0, imgData[2][i][j]))
			newImg.p[i][j].a = img.p[i][j].a
		}, img,
	)
}

// setValuesZero sets values in the Haar matrix to zero based on the compression ratio
func setValuesZero(comp [][]float64, compRatio float64) {
	if compRatio == 0 {
		return
	}

	flattened := []float64{}
	for _, row := range comp {
		for _, val := range row {
			if math.Abs(val) >= 0.001 {
				flattened = append(flattened, math.Abs(val))
			}
		}
	}

	sort.Float64s(flattened)
	threshold := flattened[int(compRatio*float64(len(flattened)-1))]

	for i := range comp {
		for j := range comp[i] {
			if math.Abs(comp[i][j]) <= threshold {
				comp[i][j] = 0
			}
		}
	}
}

// haarTransform2d performs the Haar transform on a 2D matrix
func haarTransform2d(comp [][]float64, size int) {
	c := size
	for c > 1 {
		len := c
		for i := 0; i < size; i++ {
			haarPartial(comp[i], len)
		}
		for i := 0; i < size; i++ {
			col := make([]float64, len)
			for j := 0; j < len; j++ {
				col[j] = comp[j][i]
			}
			haarPartial(col, len)
			for j := 0; j < len; j++ {
				comp[j][i] = col[j]
			}
		}
		c /= 2
	}
}

// inverseHaarTransform2d performs the inverse Haar transform on a 2D matrix
func inverseHaarTransform2d(comp [][]float64, size int) {
	c := 2
	for c <= size {
		len := c
		for i := 0; i < size; i++ {
			col := make([]float64, len)
			for j := 0; j < len; j++ {
				col[j] = comp[j][i]
			}
			inverseHaarPartial(col, len)
			for j := 0; j < len; j++ {
				comp[j][i] = col[j]
			}
		}
		for i := 0; i < size; i++ {
			inverseHaarPartial(comp[i], len)
		}
		c *= 2
	}
}

// haarPartial performs the Haar transform on a single row or column
func haarPartial(comp []float64, length int) {
	sqrt2 := math.Sqrt(2)
	pass := make([]float64, length)
	for i := 0; i < length; i += 2 {
		pass[i/2] = (comp[i] + comp[i+1]) / sqrt2
		pass[(i+length)/2] = (comp[i] - comp[i+1]) / sqrt2
	}
	copy(comp, pass)
}

// inverseHaarPartial performs the inverse Haar transform on a single row or column
func inverseHaarPartial(comp []float64, length int) {
	sqrt2 := math.Sqrt(2)
	pass := make([]float64, length)
	for i := 0; i < length; i += 2 {
		pass[i] = (comp[i/2] + comp[(i+length)/2]) / sqrt2
		pass[i+1] = (comp[i/2] - comp[(i+length)/2]) / sqrt2
	}
	copy(comp, pass)
}

// calculate the next power of 2 greater than num
func nextPerfectPowerOf2(num int) int {
	res := 1
	for res < num {
		res <<= 1
	}
	return res
}
