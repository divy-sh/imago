package image

import (
	"math"
	"testing"
)

func TestBlur(t *testing.T) {
	img, _ := NewImage(1, 1)
	img.p[0][0] = Pixel{r: 0.4, g: 0.2, b: 0.3, a: 1}

	img, _ = img.Blur()

	if math.Abs(img.p[0][0].r-0.1) >= 0.001 || math.Abs(img.p[0][0].g-0.05) >= 0.001 || math.Abs(img.p[0][0].b-0.075) >= 0.001 {
		t.Error("GetGrayScaleByIntensity function failed")
	}
}

func TestSharpen(t *testing.T) {
	img, _ := NewImage(1, 1)
	img.p[0][0] = Pixel{r: 0.4, g: 0.2, b: 0.3, a: 1}

	img, _ = img.Sharpen()

	if math.Abs(img.p[0][0].r-0.4) >= 0.001 || math.Abs(img.p[0][0].g-0.2) >= 0.001 || math.Abs(img.p[0][0].b-0.3) >= 0.001 {
		t.Error("GetGrayScaleByIntensity function failed")
	}
}

func TestEdgeDetection(t *testing.T) {
	img, _ := NewImage(1, 1)
	img.p[0][0] = Pixel{r: 0.4, g: 0.2, b: 0.3, a: 1}

	img, _ = img.EdgeDetect()

	if img.p[0][0].r != 1 || img.p[0][0].g != 1 || img.p[0][0].b != 1 {
		t.Error("EdgeDetect function failed")
	}
}
