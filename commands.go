package main

import (
	"errors"
	"strconv"

	"github.com/divy-sh/imago/image"
)

type Command struct {
	Func          func(*image.Img, []string) (*image.Img, error)
	Documentation string
}

var commands = map[string]Command{
	"horizontalFlip": {
		Func: horizontalFlip,
		Documentation: `horizontalFlip <input path> <output path>
		Flip the image horizontally.`,
	},
	"verticalFlip": {
		Func: verticalFlip,
		Documentation: `verticalFlip <input path> <output path>
		Flip the image vertically.`,
	},
	"brighten": {
		Func: brighten,
		Documentation: `brighten <input path> <output path> <brightness factor, between (-100, 100)>
		Adjust the brightness of the image.`,
	},
	"getRed": {
		Func: getRed,
		Documentation: `getRed <input path> <output path>
		Extract specific color components (red) from the image.`,
	},
	"getGreen": {
		Func: getGreen,
		Documentation: `getGreen <input path> <output path>
		Extract specific color components (green) from the image.`,
	},
	"getBlue": {
		Func: getBlue,
		Documentation: `getBlue <input path> <output path>
		Extract specific color components (blue) from the image.`,
	},
	"getGrayScaleByValue": {
		Func: getGrayScaleByValue,
		Documentation: `getGrayScaleByValue <input path> <output path>
		Convert the image to greyscale using the value component of the pixels.`,
	},
	"getGrayScaleByIntensity": {
		Func: getGrayScaleByIntensity,
		Documentation: `getGrayScaleByIntensity <input path> <output path>
		Convert the image to greyscale using the intensity component of the pixels.`,
	},
	"haarCompress": {
		Func: haarCompress,
		Documentation: `haarCompress <input path> <output path> <compression ratio, between (0, 1)>
		Compress the image using the Haar wavelet transform.`,
	},
	"sharpen": {
		Func: sharpen,
		Documentation: `sharpen <input path> <output path>
		Apply a sharpen filter to the image.`,
	},
	"blur": {
		Func: blur,
		Documentation: `blur <input path> <output path>
		Apply a blur filter to the image.`,
	},
	"edgedetect": {
		Func: edgedetect,
		Documentation: `edgedetect <input path> <output path>
		Apply an edge detection filter to the image.`,
	},
	"getHistogram": {
		Func: getHistogram,
		Documentation: `getHistogram <input path> <output path> <histogram size>
		Generate a histogram of the image.`,
	},
	"colorCorrect": {
		Func: colorCorrect,
		Documentation: `colorCorrect <input path> <output path>
		Apply color correction to the image.`,
	},
}

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
		return nil, errors.New("compression ratio not provided\n usage: haarCompress <input path> <output path> <compression ratio")
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

func getHistogram(img *image.Img, args []string) (*image.Img, error) {
	if len(args) < 1 {
		return nil, errors.New("histogram size not provided\n usage: getHistogram <input path> <output path> <histogram size>")
	}
	histSize, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return nil, errors.New("invalid histogram size")
	}
	return img.GetHistogram(int(histSize))
}

func colorCorrect(img *image.Img, _ []string) (*image.Img, error) {
	return img.ColorCorrect()
}
