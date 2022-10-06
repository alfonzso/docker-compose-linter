package main

import (
	// "encoding/json"
	// "flag"
	"fmt"
	"os"
	// "strings"

	"github.com/alfonzso/docker-compose-linter/linter"
)

func main() {
	// appDescription := "\nDocker Compose file linter: Validate your 'docker-compose.yml' file\n"
	// appUsage := fmt.Sprintf("Usage: %s -f '<filename.xyz>' [ OPTIONS ] \n", strings.Replace(os.Args[0], "./", "", -1))

	// filename := flag.String("f", "", "filename, e.x.: -f '<your-shine-compose-file>.yaml' \t [ manadotry ] ")
	// verbose := flag.Bool("v", false, "Verbose")

	// flag.Usage = func() {
	// 	fmt.Fprintf(flag.CommandLine.Output(), "\n%s %s\n", appUsage, appDescription)
	// 	fmt.Println("Options:")
	// 	flag.PrintDefaults()
	// }

	// flag.Parse()

	filename, verbose := Flags()

	if *filename == "" {
		fmt.Println("[ ERROR ] Filename missing !")
		fmt.Println()
		fmt.Println(appUsage)
		os.Exit(1)
	}

	b, err := os.ReadFile(*filename)
	if err != nil {
		fmt.Print(err)
	}

	exitcode := linter.Parser(string(b), *verbose)
	os.Exit(exitcode)
}
