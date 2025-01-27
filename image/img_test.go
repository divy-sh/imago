package image

import (
	"fmt"
	"math"
	"testing"
)

func TestNewImage(t *testing.T) {
	img, err := NewImage(10, 10)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if img.h != 10 || img.w != 10 {
		t.Errorf("Expected dimensions 10x10, got %dx%d", img.h, img.w)
	}

	_, err = NewImage(0, 10)
	if err == nil {
		t.Error("Expected error for zero height, got none")
	}

	_, err = NewImage(10, 0)
	if err == nil {
		t.Error("Expected error for zero width, got none")
	}
}

func TestHorizontalFlip(t *testing.T) {
	img, _ := NewImage(2, 2)
	img.p[0][0] = Pixel{r: 1, g: 0, b: 0, a: 1}
	img.p[0][1] = Pixel{r: 0, g: 1, b: 0, a: 1}
	img.p[1][0] = Pixel{r: 0, g: 0, b: 1, a: 1}
	img.p[1][1] = Pixel{r: 1, g: 1, b: 1, a: 1}

	flippedImg, _ := img.HorizontalFlip()

	if flippedImg.p[0][0] != img.p[0][1] || flippedImg.p[0][1] != img.p[0][0] {
		t.Error("Horizontal flip failed")
	}
}

func TestVerticalFlip(t *testing.T) {
	img, _ := NewImage(2, 2)
	img.p[0][0] = Pixel{r: 1, g: 0, b: 0, a: 1}
	img.p[1][0] = Pixel{r: 0, g: 1, b: 0, a: 1}

	flippedImg, _ := img.VerticalFlip()

	if flippedImg.p[0][0] != img.p[1][0] || flippedImg.p[1][0] != img.p[0][0] {
		t.Error("Vertical flip failed")
	}
}

func TestBrighten(t *testing.T) {
	img, _ := NewImage(1, 1)
	img.p[0][0] = Pixel{r: 0.5, g: 0.5, b: 0.5, a: 1}

	brightenedImg, _ := img.Brighten(10)

	if brightenedImg.p[0][0].r != 0.6 || brightenedImg.p[0][0].g != 0.6 || brightenedImg.p[0][0].b != 0.6 {
		t.Error("Brighten function failed")
	}
}

func TestImgBrightenInvalidValue(t *testing.T) {
	tests := []struct {
		name string
		bVal int
	}{
		{"haar compression ratio less than -100", -101},
		{"haar compression ratio greater than 100", 101},
	}
	for _, tt := range tests {
		img, _ := NewImage(1, 1)
		img.p[0][0] = Pixel{r: 0.5, g: 0.2, b: 0.3, a: 1}

		_, err := img.Brighten(float64(tt.bVal))
		if err == nil || err.Error() != "invalid brightness value" {
			t.Error("Expected error for invalid value")
		}
	}
}

func TestGetRed(t *testing.T) {
	img, _ := NewImage(1, 1)
	img.p[0][0] = Pixel{r: 0.5, g: 0.2, b: 0.3, a: 1}

	redImg, _ := img.GetRed()

	if redImg.p[0][0].r != 0.5 || redImg.p[0][0].g != 0 || redImg.p[0][0].b != 0 {
		t.Error("GetRed function failed")
	}
}

func TestGetGreen(t *testing.T) {
	img, _ := NewImage(1, 1)
	img.p[0][0] = Pixel{r: 0.5, g: 0.2, b: 0.3, a: 1}

	greenImg, _ := img.GetGreen()

	if greenImg.p[0][0].g != 0.2 || greenImg.p[0][0].r != 0 || greenImg.p[0][0].b != 0 {
		t.Error("GetGreen function failed")
	}
}

func TestGetBlue(t *testing.T) {
	img, _ := NewImage(1, 1)
	img.p[0][0] = Pixel{r: 0.5, g: 0.2, b: 0.3, a: 1}

	blueImg, _ := img.GetBlue()

	if blueImg.p[0][0].b != 0.3 || blueImg.p[0][0].r != 0 || blueImg.p[0][0].g != 0 {
		t.Error("GetBlue function failed")
	}
}

func TestGetGrayScaleByValue(t *testing.T) {
	img, _ := NewImage(1, 1)
	img.p[0][0] = Pixel{r: 0.5, g: 0.2, b: 0.3, a: 1}

	grayImg, _ := img.GetGrayScaleByValue()

	if grayImg.p[0][0].b != 0.5 || grayImg.p[0][0].r != 0.5 || grayImg.p[0][0].g != 0.5 {
		t.Error("GetGrayScaleByValue function failed")
	}
}

func TestImgHaarCompress(t *testing.T) {
	img, _ := NewImage(1, 1)
	img.p[0][0] = Pixel{r: 0.5, g: 0.2, b: 0.3, a: 1}

	compressed, _ := img.HaarCompress(0.5)
	if compressed.p[0][0].b != 0 || compressed.p[0][0].r != 0 || compressed.p[0][0].g != 0 {
		t.Error("haar compress function failed")
	}
}

func TestImgHaarCompressInvalidRatio(t *testing.T) {
	tests := []struct {
		name             string
		compressionRatio int
	}{
		{"haar compression ratio less than zero", -1},
		{"haar compression ratio greater than one", 2},
	}
	for _, tt := range tests {
		img, _ := NewImage(1, 1)
		img.p[0][0] = Pixel{r: 0.5, g: 0.2, b: 0.3, a: 1}

		_, err := img.HaarCompress(float64(tt.compressionRatio))
		if err == nil || err.Error() != "invalid compression ratio" {
			t.Error("Expected error for invalid ratio")
		}
	}
}

func TestGetGrayScaleByIntensity(t *testing.T) {
	img, _ := NewImage(1, 1)
	img.p[0][0] = Pixel{r: 0.4, g: 0.2, b: 0.3, a: 1}

	grayImg, _ := img.GetGrayScaleByIntensity()

	if math.Abs(grayImg.p[0][0].b-0.3) >= 0.001 || math.Abs(grayImg.p[0][0].g-0.3) >= 0.001 || math.Abs(grayImg.p[0][0].r-0.3) >= 0.001 {
		t.Error("GetGrayScaleByIntensity function failed")
	}
}

func TestColorCorrection(t *testing.T) {
	img, _ := NewImage(1, 1)
	img.p[0][0] = Pixel{r: 0.9, g: 0.2, b: 0.3, a: 1}

	correctedImg, _ := img.ColorCorrect()
	fmt.Println(correctedImg.p[0][0])
	if correctedImg.p[0][0].r != 0.9 || correctedImg.p[0][0].g != 0.2 || correctedImg.p[0][0].b != 0.3 {
		t.Error("ColorCorrection function failed")
	}
}

func TestImgGetHistogram(t *testing.T) {
	img, _ := NewImage(1, 1)
	img.p[0][0] = Pixel{r: 0.9, g: 0.2, b: 0.3, a: 1}

	hist, _ := img.GetHistogram(256)

	fmt.Println(hist.p[0][51])
	if hist.p[0][51].r != 0 || hist.p[0][51].g != 0 || hist.p[0][51].b != 1 {
		t.Error("GetHistogram function failed")
	}
}

func TestImgGetHistogramSizeTooSmall(t *testing.T) {
	img, _ := NewImage(1, 1)
	img.p[0][0] = Pixel{r: 0.9, g: 0.2, b: 0.3, a: 1}

	_, err := img.GetHistogram(99)
	fmt.Println(err)
	if err == nil || err.Error() != "histogram size too small" {
		t.Error("Expected error for small histogram size")
	}
}

func TestClampPixelValue(t *testing.T) {
	if clampPixelValue(0.5) != 0.5 {
		t.Error("ClampPixelValue failed for normal value")
	}
	if clampPixelValue(-0.1) != 0 {
		t.Error("ClampPixelValue failed for negative value")
	}
}

func TestProcess(t *testing.T) {
	img, _ := NewImage(2, 2)
	img.p[0][0] = Pixel{r: 1, g: 0, b: 0, a: 1}
	img.p[0][1] = Pixel{r: 0, g: 1, b: 0, a: 1}

	newImg, err := process(func(i, j int, newImg *Img) {
		newImg.p[i][j] = img.p[i][j]
	}, img)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if newImg.p[0][0] != img.p[0][0] || newImg.p[0][1] != img.p[0][1] {
		t.Error("Process function failed")
	}
}
