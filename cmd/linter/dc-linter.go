package main

import (
	"os"

	"github.com/alfonzso/docker-compose-linter/linter"
	"github.com/alfonzso/docker-compose-linter/pkg/flags"
	"github.com/alfonzso/docker-compose-linter/pkg/input"
)

func main() {

	filename := flags.Parser()
	composeFileContentAsString := input.Manager(filename)

	os.Exit(linter.Run(composeFileContentAsString))
}
