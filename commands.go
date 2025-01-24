package main

import (
	"errors"
	"strconv"

	"github.com/divy-sh/imago/image"
)

func horizontalFlip(img *image.Img, _ []string) (*image.Img, error) {
	return img.HorizontalFlip()
}

func verticalFlip(img *image.Img, _ []string) (*image.Img, error) {
	return img.VerticalFlip()
}

func brighten(img *image.Img, args []string) (*image.Img, error) {
	if len(args) < 1 {
		return nil, errors.New("brightness value not provided\n usage: brighten <input path> <output path> <brightness value>")
	}
	brightness, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return nil, errors.New("invalid brightness value")
	}
	return img.Brighten(brightness)
}

func getRed(img *image.Img, _ []string) (*image.Img, error) {
	return img.GetRed()
}

func getGreen(img *image.Img, _ []string) (*image.Img, error) {
	return img.GetGreen()
}

func getBlue(img *image.Img, _ []string) (*image.Img, error) {
	return img.GetBlue()
}

func getGrayScaleByValue(img *image.Img, _ []string) (*image.Img, error) {
	return img.GetGrayScaleByValue()
}

func getGrayScaleByIntensity(img *image.Img, _ []string) (*image.Img, error) {
	return img.GetGrayScaleByIntensity()
}

func haarCompress(img *image.Img, args []string) (*image.Img, error) {
	if len(args) < 1 {
		return nil, errors.New("compression ratio not provided\n usage: brighten <input path> <output path> <brightness value>")
	}
	compressionRatio, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return nil, errors.New("invalid compression ratio")
	}
	return img.HaarCompress(compressionRatio)
}

func sharpen(img *image.Img, _ []string) (*image.Img, error) {
	return img.Sharpen()
}

func blur(img *image.Img, _ []string) (*image.Img, error) {
	return img.Blur()
}

func edgedetect(img *image.Img, _ []string) (*image.Img, error) {
	return img.EdgeDetect()
}
