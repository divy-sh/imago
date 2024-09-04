package app

import (
	"math"
	"os"
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

func TestLoadAndSaveAsPNG(t *testing.T) {
	img, err := Load("../testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	err = img.SaveAsPNG("../testdata/output.png")
	if err != nil {
		t.Fatalf("Failed to save image: %v", err)
	}

	// Check if file was created
	if _, err := os.Stat("../testdata/output.png"); os.IsNotExist(err) {
		t.Errorf("Expected file output.png to exist, but it does not")
	}
	_ = os.Remove("../testdata/output.png") // Cleanup
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

	brightenedImg, _ := img.Brighten(0.1)

	if brightenedImg.p[0][0].r != 0.6 || brightenedImg.p[0][0].g != 0.6 || brightenedImg.p[0][0].b != 0.6 {
		t.Error("Brighten function failed")
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

func TestGetGrayScaleByIntensity(t *testing.T) {
	img, _ := NewImage(1, 1)
	img.p[0][0] = Pixel{r: 0.4, g: 0.2, b: 0.3, a: 1}

	grayImg, _ := img.GetGrayScaleByIntensity()

	if math.Abs(grayImg.p[0][0].b-0.3) >= 0.001 || math.Abs(grayImg.p[0][0].g-0.3) >= 0.001 || math.Abs(grayImg.p[0][0].r-0.3) >= 0.001 {
		t.Error("GetGrayScaleByIntensity function failed")
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
