package image

import (
	"testing"
)

func TestGetHistogram(t *testing.T) {
	img, err := Load("../testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	histSize := 256
	hist, err := getHistogram(img, histSize)
	if err != nil {
		t.Fatalf("Failed to get histogram: %v", err)
	}

	if hist == nil {
		t.Fatalf("Expected histogram image, got nil")
	}

	if hist.w != histSize+1 || hist.h != histSize+1 {
		t.Errorf("Expected histogram size %dx%d, got %dx%d", histSize+1, histSize+1, hist.w, hist.h)
	}
}
