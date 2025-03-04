package main

import (
	"os"
	"testing"

	"github.com/divy-sh/imago/image"
)

func TestHorizontalFlip(t *testing.T) {
	img, err := image.Load("./testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	flippedImg, err := horizontalFlip(img, nil)
	if err != nil {
		t.Fatalf("Failed to flip image horizontally: %v", err)
	}

	err = image.SaveImage(flippedImg, "./testdata/output_horizontal_flip.png")
	if err != nil {
		t.Fatalf("Failed to save flipped image: %v", err)
	}

	if _, err := os.Stat("./testdata/output_horizontal_flip.png"); os.IsNotExist(err) {
		t.Errorf("Expected file output_horizontal_flip.png to exist, but it does not")
	}
	_ = os.Remove("./testdata/output_horizontal_flip.png") // Cleanup
}

func TestVerticalFlip(t *testing.T) {
	img, err := image.Load("./testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	flippedImg, err := verticalFlip(img, nil)
	if err != nil {
		t.Fatalf("Failed to flip image vertically: %v", err)
	}

	err = image.SaveImage(flippedImg, "./testdata/output_vertical_flip.png")
	if err != nil {
		t.Fatalf("Failed to save flipped image: %v", err)
	}

	if _, err := os.Stat("./testdata/output_vertical_flip.png"); os.IsNotExist(err) {
		t.Errorf("Expected file output_vertical_flip.png to exist, but it does not")
	}
	_ = os.Remove("./testdata/output_vertical_flip.png") // Cleanup
}

func TestBrighten(t *testing.T) {
	img, err := image.Load("./testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	brightenedImg, err := brighten(img, []string{"1.5"})
	if err != nil {
		t.Fatalf("Failed to brighten image: %v", err)
	}

	err = image.SaveImage(brightenedImg, "./testdata/output_brighten.png")
	if err != nil {
		t.Fatalf("Failed to save brightened image: %v", err)
	}

	if _, err := os.Stat("./testdata/output_brighten.png"); os.IsNotExist(err) {
		t.Errorf("Expected file output_brighten.png to exist, but it does not")
	}
	_ = os.Remove("./testdata/output_brighten.png") // Cleanup
}

func TestBrightenInvalidArguments(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "No arguments",
			args: []string{},
		},
		{
			name: "Invalid argument",
			args: []string{"invalid"},
		},
	}
	for _, tt := range tests {
		img, err := image.Load("./testdata/test.png")
		if err != nil {
			t.Fatalf("Failed to load image: %v", err)
		}

		_, err = brighten(img, tt.args)
		if err == nil {
			t.Fatalf("expected failure to brighten image, go no error")
		}
	}
}

func TestGetRed(t *testing.T) {
	img, err := image.Load("./testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	redImg, err := getRed(img, nil)
	if err != nil {
		t.Fatalf("Failed to get red channel: %v", err)
	}

	err = image.SaveImage(redImg, "./testdata/output_red.png")
	if err != nil {
		t.Fatalf("Failed to save red channel image: %v", err)
	}

	if _, err := os.Stat("./testdata/output_red.png"); os.IsNotExist(err) {
		t.Errorf("Expected file output_red.png to exist, but it does not")
	}
	_ = os.Remove("./testdata/output_red.png") // Cleanup
}

func TestGetGreen(t *testing.T) {
	img, err := image.Load("./testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	greenImg, err := getGreen(img, nil)
	if err != nil {
		t.Fatalf("Failed to get green channel: %v", err)
	}

	err = image.SaveImage(greenImg, "./testdata/output_green.png")
	if err != nil {
		t.Fatalf("Failed to save green channel image: %v", err)
	}

	if _, err := os.Stat("./testdata/output_green.png"); os.IsNotExist(err) {
		t.Errorf("Expected file output_green.png to exist, but it does not")
	}
	_ = os.Remove("./testdata/output_green.png") // Cleanup
}

func TestGetBlue(t *testing.T) {
	img, err := image.Load("./testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	blueImg, err := getBlue(img, nil)
	if err != nil {
		t.Fatalf("Failed to get blue channel: %v", err)
	}

	err = image.SaveImage(blueImg, "./testdata/output_blue.png")
	if err != nil {
		t.Fatalf("Failed to save blue channel image: %v", err)
	}

	if _, err := os.Stat("./testdata/output_blue.png"); os.IsNotExist(err) {
		t.Errorf("Expected file output_blue.png to exist, but it does not")
	}
	_ = os.Remove("./testdata/output_blue.png") // Cleanup
}

func TestGetGrayScaleByValue(t *testing.T) {
	img, err := image.Load("./testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	grayImg, err := getGrayScaleByValue(img, nil)
	if err != nil {
		t.Fatalf("Failed to get grayscale by value: %v", err)
	}

	err = image.SaveImage(grayImg, "./testdata/output_gray_value.png")
	if err != nil {
		t.Fatalf("Failed to save grayscale by value image: %v", err)
	}

	if _, err := os.Stat("./testdata/output_gray_value.png"); os.IsNotExist(err) {
		t.Errorf("Expected file output_gray_value.png to exist, but it does not")
	}
	_ = os.Remove("./testdata/output_gray_value.png") // Cleanup
}

func TestGetGrayScaleByIntensity(t *testing.T) {
	img, err := image.Load("./testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	grayImg, err := getGrayScaleByIntensity(img, nil)
	if err != nil {
		t.Fatalf("Failed to get grayscale by intensity: %v", err)
	}

	err = image.SaveImage(grayImg, "./testdata/output_gray_intensity.png")
	if err != nil {
		t.Fatalf("Failed to save grayscale by intensity image: %v", err)
	}

	if _, err := os.Stat("./testdata/output_gray_intensity.png"); os.IsNotExist(err) {
		t.Errorf("Expected file output_gray_intensity.png to exist, but it does not")
	}
	_ = os.Remove("./testdata/output_gray_intensity.png") // Cleanup
}

func TestHaarCompress(t *testing.T) {
	img, err := image.Load("./testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	compressedImg, err := haarCompress(img, []string{"0.5"})
	if err != nil {
		t.Fatalf("Failed to compress image: %v", err)
	}

	err = image.SaveImage(compressedImg, "./testdata/output_haar_compress.png")
	if err != nil {
		t.Fatalf("Failed to save compressed image: %v", err)
	}

	if _, err := os.Stat("./testdata/output_haar_compress.png"); os.IsNotExist(err) {
		t.Errorf("Expected file output_haar_compress.png to exist, but it does not")
	}
	_ = os.Remove("./testdata/output_haar_compress.png") // Cleanup
}

func TestHaarCompressInvalidArguments(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "No arguments",
			args: []string{},
		},
		{
			name: "Invalid argument",
			args: []string{"invalid"},
		},
	}
	for _, tt := range tests {
		img, err := image.Load("./testdata/test.png")
		if err != nil {
			t.Fatalf("Failed to load image: %v", err)
		}

		_, err = haarCompress(img, tt.args)
		if err == nil {
			t.Fatalf("expected failure to compress image, go no error")
		}
	}
}

func TestSharpen(t *testing.T) {
	img, err := image.Load("./testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	sharpenedImg, err := sharpen(img, nil)
	if err != nil {
		t.Fatalf("Failed to sharpen image: %v", err)
	}

	err = image.SaveImage(sharpenedImg, "./testdata/output_sharpen.png")
	if err != nil {
		t.Fatalf("Failed to save sharpened image: %v", err)
	}

	if _, err := os.Stat("./testdata/output_sharpen.png"); os.IsNotExist(err) {
		t.Errorf("Expected file output_sharpen.png to exist, but it does not")
	}
	_ = os.Remove("./testdata/output_sharpen.png") // Cleanup
}

func TestBlur(t *testing.T) {
	img, err := image.Load("./testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	blurredImg, err := blur(img, nil)
	if err != nil {
		t.Fatalf("Failed to blur image: %v", err)
	}

	err = image.SaveImage(blurredImg, "./testdata/output_blur.png")
	if err != nil {
		t.Fatalf("Failed to save blurred image: %v", err)
	}

	if _, err := os.Stat("./testdata/output_blur.png"); os.IsNotExist(err) {
		t.Errorf("Expected file output_blur.png to exist, but it does not")
	}
	_ = os.Remove("./testdata/output_blur.png") // Cleanup
}

func TestEdgeDetect(t *testing.T) {
	img, err := image.Load("./testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	edgeImg, err := edgedetect(img, nil)
	if err != nil {
		t.Fatalf("Failed to detect edges: %v", err)
	}

	err = image.SaveImage(edgeImg, "./testdata/output_edge_detect.png")
	if err != nil {
		t.Fatalf("Failed to save edge detected image: %v", err)
	}

	if _, err := os.Stat("./testdata/output_edge_detect.png"); os.IsNotExist(err) {
		t.Errorf("Expected file output_edge_detect.png to exist, but it does not")
	}
	_ = os.Remove("./testdata/output_edge_detect.png") // Cleanup
}

func TestDrawHistogram(t *testing.T) {
	img, err := image.Load("./testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	edgeImg, err := getHistogram(img, []string{"255"})
	if err != nil {
		t.Fatalf("Failed to get histogram: %v", err)
	}

	err = image.SaveImage(edgeImg, "./testdata/output_histogram.png")
	if err != nil {
		t.Fatalf("Failed to save histogram image: %v", err)
	}

	if _, err := os.Stat("./testdata/output_histogram.png"); os.IsNotExist(err) {
		t.Errorf("Expected file output_histogram.png to exist, but it does not")
	}
	_ = os.Remove("./testdata/output_histogram.png") // Cleanup
}

func TestGetHistogramInvalidArguments(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "No arguments",
			args: []string{},
		},
		{
			name: "Invalid argument",
			args: []string{"invalid"},
		},
	}
	for _, tt := range tests {
		img, err := image.Load("./testdata/test.png")
		if err != nil {
			t.Fatalf("Failed to load image: %v", err)
		}

		_, err = getHistogram(img, tt.args)
		if err == nil {
			t.Fatalf("expected failure to brighten image, go no error")
		}
	}
}

func TestColorCorrect(t *testing.T) {
	img, err := image.Load("./testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	edgeImg, err := colorCorrect(img, []string{})
	if err != nil {
		t.Fatalf("Failed to color correct: %v", err)
	}

	err = image.SaveImage(edgeImg, "./testdata/output_color_correct.png")
	if err != nil {
		t.Fatalf("Failed to save color corrected image: %v", err)
	}

	if _, err := os.Stat("./testdata/output_color_correct.png"); os.IsNotExist(err) {
		t.Errorf("Expected file output_color_correct.png to exist, but it does not")
	}
	_ = os.Remove("./testdata/output_color_correct.png") // Cleanup
}

func TestLevelAdjust(t *testing.T) {
	OutputPath := "./testdata/output_level_adjusted.png"
	img, err := image.Load("./testdata/test.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	levelAdjustedImg, err := levelAdjust(img, []string{"0.1", "0.3", "0.7"})
	if err != nil {
		t.Fatalf("Failed to color correct: %v", err)
	}

	err = image.SaveImage(levelAdjustedImg, OutputPath)
	if err != nil {
		t.Fatalf("Failed to save level adjusted image: %v", err)
	}

	if _, err := os.Stat(OutputPath); os.IsNotExist(err) {
		t.Errorf("Expected file %s to exist, but it does not", OutputPath)
	}
	_ = os.Remove(OutputPath) // Cleanup
}
