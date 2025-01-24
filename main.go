package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	err := execute(args)
	if err != nil {
		fmt.Println(err)
	}
}
