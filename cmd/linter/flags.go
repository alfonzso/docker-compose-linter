package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var appDescription = "\nDocker Compose file linter: Validate your 'docker-compose.yml' file\n"
var appUsage = fmt.Sprintf("Usage: %s -f '<filename.xyz>' [ OPTIONS ] \n", strings.Replace(os.Args[0], "./", "", -1))

func Flags() (string, bool) {

	filename := flag.String("f", "", "filename, e.x.: -f '<your-shine-compose-file>.yaml' \t [ manadotry ] ")
	verbose := flag.Bool("v", false, "Verbose")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "\n%s %s\n", appUsage, appDescription)
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	return *filename, *verbose
}
