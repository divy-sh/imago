package app

import (
	"os"
	"testing"
)

func TestLoadAndSaveAsPng(t *testing.T) {
	img, err := Load("../testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	err = SaveImage(img, "../testdata/output.png")
	if err != nil {
		t.Fatalf("Failed to save image: %v", err)
	}

	// Check if file was created
	if _, err := os.Stat("../testdata/output.png"); os.IsNotExist(err) {
		t.Errorf("Expected file output.png to exist, but it does not")
	}
	_ = os.Remove("../testdata/output.png") // Cleanup
}

func TestLoadAndSaveAsJpeg(t *testing.T) {
	img, err := Load("../testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	err = SaveImage(img, "../testdata/output.jpeg")
	if err != nil {
		t.Fatalf("Failed to save image: %v", err)
	}

	// Check if file was created
	if _, err := os.Stat("../testdata/output.jpeg"); os.IsNotExist(err) {
		t.Errorf("Expected file output.jpeg to exist, but it does not")
	}
	_ = os.Remove("../testdata/output.jpeg") // Cleanup
}
