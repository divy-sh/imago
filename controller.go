package main

import (
	"errors"

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
}

func execute(args []string) error {
	if len(args) == 1 && args[0] == "help" {
		help()
		return nil
	}
	if len(args) < 3 {
		return errors.New(`incorrect command provided. usage: <command> <input path> <output path> <other args>`)
	}
	command := args[0]
	inputPath := args[1]
	outputPath := args[2]
	args = args[3:]

	img, err := image.Load(inputPath)
	if err != nil {
		return err
	}

	img, err = commands[command].Func(img, args)
	if err != nil {
		return err
	}
	err = image.SaveImage(img, outputPath)
	if err != nil {
		return err
	}
	return nil
}

func help() {
	for _, cmd := range commands {
		println(cmd.Documentation, "\n")
	}
}
