package main

import (
	"fmt"
	"os"

	"github.com/alfonzso/docker-compose-linter/linter"
	"github.com/alfonzso/docker-compose-linter/pkg/flags"
)

func main() {

	filename, verbose := flags.Flags()

	if filename == "" {
		fmt.Println("[ ERROR ] Filename missing !")
		fmt.Println()
		fmt.Println(flags.AppUsage)
		os.Exit(1)
	}

	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	os.Exit(linter.Run(string(b), verbose))
}
