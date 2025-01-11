package app

import (
	"fmt"
	"testing"
)

func TestHaarCompress(t *testing.T) {
	img, _ := NewImage(4, 4)
	img.p[0][0] = Pixel{r: 1, g: 0, b: 0, a: 1}
	img.p[0][1] = Pixel{r: 0, g: 1, b: 0, a: 1}
	img.p[0][2] = Pixel{r: 0, g: 0, b: 1, a: 1}
	img.p[0][3] = Pixel{r: 1, g: 1, b: 1, a: 1}
	img.p[1][0] = Pixel{r: 1, g: 0, b: 0, a: 1}
	img.p[1][1] = Pixel{r: 0, g: 1, b: 0, a: 1}
	img.p[1][2] = Pixel{r: 0, g: 0, b: 1, a: 1}
	img.p[1][3] = Pixel{r: 1, g: 1, b: 1, a: 1}
	img.p[2][0] = Pixel{r: 1, g: 0, b: 0, a: 1}
	img.p[2][1] = Pixel{r: 0, g: 1, b: 0, a: 1}
	img.p[2][2] = Pixel{r: 0, g: 0, b: 1, a: 1}
	img.p[2][3] = Pixel{r: 1, g: 1, b: 1, a: 1}
	img.p[3][0] = Pixel{r: 1, g: 0, b: 0, a: 1}
	img.p[3][1] = Pixel{r: 0, g: 1, b: 0, a: 1}
	img.p[3][2] = Pixel{r: 0, g: 0, b: 1, a: 1}
	img.p[3][3] = Pixel{r: 1, g: 1, b: 1, a: 1}

	tests := []struct {
		ratio float64
	}{
		{ratio: 0.0},
		{ratio: 0.5},
		{ratio: 1.0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Ratio_%f", tt.ratio), func(t *testing.T) {
			compressedImg, err := haarCompress(img, tt.ratio)
			if err != nil {
				t.Fatalf("haarCompress failed: %v", err)
			}

			// Check if the compressed image has the same dimensions
			if compressedImg.h != img.h || compressedImg.w != img.w {
				t.Errorf("Expected dimensions %dx%d, got %dx%d", img.h, img.w, compressedImg.h, compressedImg.w)
			}

			// Check if the pixel values are within the expected range
			for i := 0; i < compressedImg.h; i++ {
				for j := 0; j < compressedImg.w; j++ {
					if compressedImg.p[i][j].r < 0 || compressedImg.p[i][j].r > 1 {
						t.Errorf("Pixel value out of range at (%d, %d): %f", i, j, compressedImg.p[i][j].r)
					}
					if compressedImg.p[i][j].g < 0 || compressedImg.p[i][j].g > 1 {
						t.Errorf("Pixel value out of range at (%d, %d): %f", i, j, compressedImg.p[i][j].g)
					}
					if compressedImg.p[i][j].b < 0 || compressedImg.p[i][j].b > 1 {
						t.Errorf("Pixel value out of range at (%d, %d): %f", i, j, compressedImg.p[i][j].b)
					}
				}
			}
		})
	}
}
