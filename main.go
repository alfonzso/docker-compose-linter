package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/docker/cli/cli/compose/loader"
	compose "github.com/docker/cli/cli/compose/types"
)

func main() {
	filename := flag.String("f", "", "docker-compose.yaml file to check \t [ manadotry ] ")
	verbose := flag.Bool("v", false, "Verbose")
	flag.Parse()

	if *filename == "" {
		fmt.Println("File needed for check!")
    fmt.Println()
    flag.Usage()
		os.Exit(1)
	}

	b, err := os.ReadFile(*filename)
	if err != nil {
		fmt.Print(err)
	}

	exitcode := parser(string(b), *verbose)
	os.Exit(exitcode)
}

func buildConfigDetails(source map[string]interface{}, env map[string]string) compose.ConfigDetails {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return compose.ConfigDetails{
		WorkingDir: workingDir,
		ConfigFiles: []compose.ConfigFile{
			{Filename: "filename.yml", Config: source},
		},
		Environment: env,
	}
}

func loadYAML(yaml string) (*compose.Config, error) {
	return loadYAMLWithEnv(yaml, nil)
}

func loadYAMLWithEnv(yaml string, env map[string]string) (*compose.Config, error) {
	dict, err := loader.ParseYAML([]byte(yaml))
	if err != nil {
		return nil, err
	}

	return loader.Load(buildConfigDetails(dict, env))
}

func parser(yaml string, verbose bool) int {
	dict, parseErr := loadYAML(yaml)

	b, err := json.Marshal(dict)
	if err != nil {
		return 1
	}

	if verbose {
		fmt.Print("Yaml: ")
		fmt.Println(string(b))
	}

	if parseErr != nil {
		fmt.Print("Error: ")
		fmt.Println(parseErr.Error())
		return 1
	}

	return 0
}
