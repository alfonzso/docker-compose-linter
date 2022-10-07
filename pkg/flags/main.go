package flags

import (
	"flag"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

var AppDescription = "\nDocker Compose file linter: Validate your 'docker-compose.yml' file\n"
var AppName = strings.Replace(os.Args[0], "./", "", -1)
var AppUsage = fmt.Sprintf(
`Usage:
    %s -f '<filename.xyz>' [ OPTIONS ]
    echo "lorem ipsum" | %s [ OPTIONS ]
`, AppName, AppName,
)

func Parser() (filename string) {

	var verbose bool
	flag.StringVar(&filename, "f", "", "fafa")
	flag.BoolVar(&verbose, "v", false, "Verbose")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "\n%s %s\n", AppUsage, AppDescription)
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	return filename
}
