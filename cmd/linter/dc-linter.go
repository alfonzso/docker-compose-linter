package main

// dostackComposeFileValidator
// dSCoFiValidator
// dcfValidator
import (
	// "encoding/json"
	"alfonzso/docker-compose-linter/linter"
	"flag"
	"fmt"
	"os"
	"strings"
	// "github.com/docker/cli/cli/compose/loader"
	// composeType "github.com/docker/cli/cli/compose/types"
)

func main() {
	appDescription := "\nDocker Compose file linter: Validate your 'docker-compose.yml' file\n"
	appUsage := fmt.Sprintf("Usage: %s -f '<filename.xyz>' [ OPTIONS ] \n", strings.Replace(os.Args[0], "./", "", -1))

	filename := flag.String("f", "", "filename, e.x.: -f '<your-shine-compose-file>.yaml' \t [ manadotry ] ")
	verbose := flag.Bool("v", false, "Verbose")

	flag.Usage = func() {
		// fmt.Fprintf(flag.CommandLine.Output(), appDescription+exampleCall+"Usage: %s -f 'filename.xyz' \n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "\n%s %s\n", appUsage, appDescription)
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *filename == "" {
		fmt.Println("[ ERROR ] Filename missing !")
		fmt.Println()
		flag.PrintDefaults()
		// flag.Usage()
		os.Exit(1)
	}

	b, err := os.ReadFile(*filename)
	if err != nil {
		fmt.Print(err)
	}

	exitcode := linter.Parser(string(b), *verbose)
	os.Exit(exitcode)
}
