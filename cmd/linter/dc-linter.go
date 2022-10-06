package main

import (
	"fmt"
	"os"

	"github.com/alfonzso/docker-compose-linter/linter"
)

func main() {

	filename, verbose := Flags()

	if filename == "" {
		fmt.Println("[ ERROR ] Filename missing !")
		fmt.Println()
		fmt.Println(appUsage)
		os.Exit(1)
	}

	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	os.Exit(linter.Run(string(b), verbose))
}
