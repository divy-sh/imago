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
		Func:          horizontalFlip,
		Documentation: "horizontalFlip <input path> <output path>",
	},
	"verticalFlip": {
		Func:          verticalFlip,
		Documentation: "verticalFlip <input path> <output path>",
	},
	"brighten": {
		Func:          brighten,
		Documentation: "brighten <input path> <output path> <brightness factor, between (-100, 100)>",
	},
	"getRed": {
		Func:          getRed,
		Documentation: "getRed <input path> <output path>",
	},
	"getGreen": {
		Func:          getGreen,
		Documentation: "getGreen <input path> <output path>",
	},
	"getBlue": {
		Func:          getBlue,
		Documentation: "getBlue <input path> <output path>",
	},
	"getGrayScaleByValue": {
		Func:          getGrayScaleByValue,
		Documentation: "getGrayScaleByValue <input path> <output path>",
	},
	"getGrayScaleByIntensity": {
		Func:          getGrayScaleByIntensity,
		Documentation: "getGrayScaleByIntensity <input path> <output path>",
	},
	"haarCompress": {
		Func:          haarCompress,
		Documentation: "haarCompress <input path> <output path> <compression ratio, between (0, 1)>",
	},
	"sharpen": {
		Func:          sharpen,
		Documentation: "sharpen <input path> <output path>",
	},
	"blur": {
		Func:          blur,
		Documentation: "blur <input path> <output path>",
	},
	"edgedetect": {
		Func:          edgedetect,
		Documentation: "edgedetect <input path> <output path>",
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
