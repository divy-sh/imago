package main

import (
	"errors"

	"github.com/divy-sh/imago/image"
)

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
