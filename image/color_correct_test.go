package image

import (
	"fmt"
	"testing"
)

func TestColorCorrect(t *testing.T) {
	// Mock image with known pixel values
	img := &Img{
		p: [][]Pixel{
			{{r: 0.3, g: 0.5, b: 0.7, a: 1}},
			{{r: 0.5, g: 0.6, b: 0.8, a: 1}},
		},
		h: 2,
		w: 1,
	}

	// Expected corrected pixel values
	expected := &Img{
		p: [][]Pixel{
			{{r: 0.3, g: 0.5, b: 0.7, a: 1}},
			{{r: 0.5, g: 0.6, b: 0.8, a: 1}},
		},
	}

	// Call colorCorrect
	correctedImg, err := colorCorrect(img)
	fmt.Println(correctedImg)
	if err != nil {
		t.Fatalf("colorCorrect returned an error: %v", err)
	}

	// Verify the output image has the expected corrected pixel values
	for i := range correctedImg.p {
		for j := range correctedImg.p[i] {
			if correctedImg.p[i][j] != expected.p[i][j] {
				t.Errorf("Pixel at (%d, %d) = %+v, want %+v", i, j, correctedImg.p[i][j], expected.p[i][j])
			}
		}
	}
}
