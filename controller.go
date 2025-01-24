package main

import (
	"errors"
	"strconv"

	"github.com/divy-sh/imago/image"
)

func execute(args []string) error {
	if len(args) < 3 {
		return errors.New("incorrect command provided\n usage: <command> <input path> <output path> <other args>")
	}
	command := args[0]
	inputPath := args[1]
	outputPath := args[2]
	args = args[3:]

	img, err := image.Load(inputPath)
	if err != nil {
		return err
	}

	img, err = callCommmand(command, img, args)
	if err != nil {
		return err
	}
	err = image.SaveImage(img, outputPath)
	if err != nil {
		return err
	}
	return nil
}

func callCommmand(command string, img *image.Img, args []string) (*image.Img, error) {
	switch command {
	case "horizontalFlip":
		return horizontalFlip(img)
	case "verticalFlip":
		return verticalFlip(img)
	case "brighten":
		return brighten(img, args)
	case "getRed":
		return getRed(img)
	case "getGreen":
		return getGreen(img)
	case "getBlue":
		return getBlue(img)
	case "getGrayScaleByValue":
		return getGrayScaleByValue(img)
	case "getGrayScaleByIntensity":
		return getGrayScaleByIntensity(img)
	case "haarCompress":
		return haarCompress(img, args)
	case "sharpen":
		return sharpen(img)
	case "blur":
		return blur(img)
	case "edgedetect":
		return edgedetect(img)
	default:
		return nil, errors.New("invalid command")
	}
}

func horizontalFlip(img *image.Img) (*image.Img, error) {
	return img.HorizontalFlip()
}

func verticalFlip(img *image.Img) (*image.Img, error) {
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

func getRed(img *image.Img) (*image.Img, error) {
	return img.GetRed()
}

func getGreen(img *image.Img) (*image.Img, error) {
	return img.GetGreen()
}

func getBlue(img *image.Img) (*image.Img, error) {
	return img.GetBlue()
}

func getGrayScaleByValue(img *image.Img) (*image.Img, error) {
	return img.GetGrayScaleByValue()
}

func getGrayScaleByIntensity(img *image.Img) (*image.Img, error) {
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

func sharpen(img *image.Img) (*image.Img, error) {
	return img.Sharpen()
}

func blur(img *image.Img) (*image.Img, error) {
	return img.Blur()
}

func edgedetect(img *image.Img) (*image.Img, error) {
	return img.EdgeDetect()
}
