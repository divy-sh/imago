package image

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadInvalidPath(t *testing.T) {
	_, err := Load("invalid/path")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestSaveNoExtension(t *testing.T) {
	img, _ := NewImage(1, 1)
	err := SaveImage(img, "output")
	fmt.Println(err)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestSaveUnsupportedExtension(t *testing.T) {
	img, _ := NewImage(1, 1)
	err := SaveImage(img, "output.gif")
	fmt.Println(err)
	if err == nil || err.Error() != "unsupported file extension: gif" {
		t.Errorf("Expected error, got nil")
	}
}

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

func TestLoadAndSaveAsJpg(t *testing.T) {
	img, err := Load("../testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	err = SaveImage(img, "../testdata/output.jpg")
	if err != nil {
		t.Fatalf("Failed to save image: %v", err)
	}

	// Check if file was created
	if _, err := os.Stat("../testdata/output.jpg"); os.IsNotExist(err) {
		t.Errorf("Expected file output.jpg to exist, but it does not")
	}
	_ = os.Remove("../testdata/output.jpg") // Cleanup
}
